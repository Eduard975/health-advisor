import google.generativeai as genai
from langchain_community.vectorstores import FAISS
from langchain_huggingface import HuggingFaceEmbeddings
from dotenv import load_dotenv
import os

load_dotenv()

# Embeddings
embeddings = HuggingFaceEmbeddings(model_name="sentence-transformers/all-MiniLM-L6-v2")

# Load vector stores
food_db = FAISS.load_local("vectors/food", embeddings, allow_dangerous_deserialization=True)
activity_db = FAISS.load_local("vectors/activity", embeddings, allow_dangerous_deserialization=True)

# Configure Gemini
genai.configure(api_key=os.getenv('GOOGLE_API_KEY'))
model = genai.GenerativeModel('gemini-2.0-flash-exp')

food_keywords = [
    "food", "nutrition", "protein", "carbs", "fats", "vitamins", "minerals",
    "diet", "weight", "calories", "meal", "snack", "breakfast", "lunch", 
    "dinner", "recipes", "ingredients", "healthy", "macro", "micro", "supplements", 
    "fiber", "sugar", "cholesterol", "hydration"
]

activity_keywords = [
    "exercise", "calories burned", "cycling", "walking", "running", "jogging",
    "weight", "workout", "sport", "lbs", "activities", "training", "strength", 
    "cardio", "flexibility", "endurance", "HIIT", "yoga", "pilates", "swimming", 
    "gym", "stretching", "steps", "movement", "fitness"
]

def get_relevant_context(db, query, k=3):
    """Retrieve top-k relevant documents from a vector store"""
    docs = db.similarity_search(query, k=k)
    return "\n\n".join([doc.page_content for doc in docs])

def classify_query(query):
    query_lower = query.lower()
    is_activity = any(word in query_lower for word in activity_keywords)
    is_food = any(word in query_lower for word in food_keywords)

    if is_food and is_activity:
        return "both"
    elif is_food:
        return "food"
    elif is_activity:
        return "activity"
    else:
        return "no"
    

def compute_dynamic_k(query, min_k=3, max_k=15):
    query_lower = query.lower()

    food_score = sum(query_lower.count(word) for word in food_keywords)
    activity_score = sum(query_lower.count(word) for word in activity_keywords)

    total = food_score + activity_score

    if total == 0:
        return min_k, min_k

    # Scale proportionally to scores
    food_k = max(min_k, min(max_k, round((food_score / total) * max_k)))
    activity_k = max(min_k, min(max_k, round((activity_score / total) * max_k)))

    return food_k, activity_k

def send_query(query):
    dataset_type = classify_query(query)
    if dataset_type == "no":
        return "I can't answer that question."

    food_k, activity_k = compute_dynamic_k(query)

    contexts = []
    if dataset_type in ["food", "both"]:
        food_context = get_relevant_context(food_db, query, k=food_k)
        if food_context:
            contexts.append(f"--- FOOD DATA ---\n{food_context}")

    if dataset_type in ["activity", "both"]:
        activity_context = get_relevant_context(activity_db, query, k=activity_k)
        if activity_context:
            contexts.append(f"--- ACTIVITY DATA ---\n{activity_context}")

    combined_context = "\n\n".join(contexts)
    
    prompt = f"""
        You are a knowledgeable nutritionist and fitness advisor. 
        Use the following context to answer the question accurately and concisely. 
        Do not invent information not in the context.

        Context:
        {combined_context}

        Question: {query}

        Instructions:
        1. Keep the answer short and structured with clear headings.
        2. Use this format:
        - Summary: 1-2 sentences about the main point.
        - Recommendations: concise actionable tips (bullet points).
        - Always end with this disclaimer: 
        \"This information is for general knowledge only 
        and does not constitute medical advice. Consult with a 
        qualified healthcare professional or registered dietitian for personalized advice.\"
        3. Provide simple metrics or calculations if needed 
        but avoid very detailed calculations or overly long explanations.
        4. If information is insufficient, clearly state it.

        Provide the answer now:
        """
    
    response = model.generate_content(prompt)
    return response.text

