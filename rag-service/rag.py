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
    "food", "foods", "nutrition", "protein", "carbs", "fats", 
    "vitamin", "vitamins", 
    "mineral", "minerals", 
    "diet", "weight", "calorie", "calories", "meal", "snack", 
    "breakfast", "lunch", "dinner", 
    "recipe", "recipes", "ingredient", "ingredients",
    "healthy", "macro", "micro", "supplement", "supplements",
    "fiber", "sugar", "cholesterol", "hydration",

    "fruit", "fruits", "vegetable", "vegetables", "meat", "fish",
    "dairy", "grain", "grains", "nut", "nuts", "seed", "seeds",
    "beverage", "drink", "eat", "eating"
]

activity_keywords = [
    "exercise", "exercises", "calorie", "calories burned",  
    "cycling", "walking", "running", "jogging",
    "weight", "workout", "workouts", "sport", "sports",
    "lbs", "lb", "kg", "activity", "activities",  
    "training", "strength", "cardio", "flexibility", 
    "endurance", "HIIT", "yoga", "pilates", "swimming", 
    "gym", "stretching", "step", "steps", "movement", "fitness"
]

def get_relevant_context(db, query, k=3):
    """Retrieve top-k relevant documents from a vector store"""
    docs = db.similarity_search(query, k=k)
    return "\n\n".join([doc.page_content for doc in docs])

def classify_query(query):
    query_lower = query.lower()
    # print(query_lower)

    is_activity = any(word in query_lower for word in activity_keywords)
    is_food = any(word in query_lower for word in food_keywords)
    
    # print(is_activity, is_food)
    if is_food and is_activity:
        return "both"
    elif is_food:
        return "food"
    elif is_activity:
        return "activity"
    else:
        return "both"
    

def compute_dynamic_k(query, min_k=5, max_k=15):
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

def send_query(query, history=None):
    if history is None:
        history = []

    chat_context = "\n".join(
        [f"{msg.sender.upper()}: {msg.text}" for msg in history[-6:]]
    )

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
        ### ROLE
        You are an expert nutritionist and fitness advisor. Your tone must be professional, helpful, and strictly evidence-based.

        ### TASK
        Your task is to analyze the provided CONTEXT to answer the USER'S QUESTION accurately.

        ### RULES
        1.  **Strict Context Adherence:** Your entire response MUST be derived exclusively from the provided CONTEXT. 
        DO NOT invent, infer, or use any external knowledge.
        2.  **Handling Insufficient Information:** If the CONTEXT does not contain enough information to answer the question, 
        you MUST respond ONLY with the following message and nothing else: "I'm sorry, but based on the provided information, 
        I cannot answer that question. Please try rephrasing or asking about a different topic."
        3.  **Conciseness:** Keep the answer structured, concise, and easy for a layperson to understand.

        ### REQUIRED OUTPUT FORMAT
        You MUST structure your response using the following markdown format exactly. Do not add any conversational text before or after this structure.

        ## Summary
        A 1-2 sentence overview of the key information from the context that answers the user's question.

        ## Recommendations
        - Use bullet points for clear, actionable advice.
        - Provide simple metrics or calculations only if they are directly supported by the context.

        ---
        ## Disclaimer
        *This information is for general knowledge only and 
        does not constitute medical advice. 
        Consult with a qualified healthcare professional or registered dietitian for personalized advice.*

        ---
        ### CHAT HISTORY CONTEXT
        {chat_context}

        ### DATASET CONTEXT
        {combined_context}

        ### USER'S QUESTION
        {query}
    """

    print('[LOG]\n' + prompt)

    response = model.generate_content(prompt)
    return response.text