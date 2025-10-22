import google.generativeai as genai
import os
from langchain_community.vectorstores import FAISS
from langchain_community.embeddings import HuggingFaceEmbeddings
from dotenv import load_dotenv

load_dotenv()

# Use HuggingFace embeddings instead of GoogleGenerativeAIEmbeddings
embeddings = HuggingFaceEmbeddings(model_name="sentence-transformers/all-MiniLM-L6-v2")
db = FAISS.load_local("data/food/food_vectorstore", embeddings, allow_dangerous_deserialization=True)

# Use Google Gemini directly
genai.configure(api_key=os.getenv('GOOGLE_API_KEY'))
model = genai.GenerativeModel('gemini-2.0-flash-exp')

def get_relevant_context(query, k=3):
    """Retrieve relevant context from the vector store"""
    docs = db.similarity_search(query, k=k)
    context = "\n\n".join([doc.page_content for doc in docs])
    return context

def rag_query(query, k=3):
    """Execute RAG query using Gemini"""
    # Get relevant context
    context = get_relevant_context(query, k=k)
    
    # Create the prompt with context
    prompt = f"""Based on the following nutritional information, answer the question:

{context}

Question: {query}

Please provide a helpful and accurate answer based on the nutritional data above:"""
    
    # Generate response
    response = model.generate_content(prompt)
    return response.text

# Your query
query = "Which foods are highest in Vitamin C and lowest in fat?"
response = rag_query(query)

print("\n🧠 Query:", query)
print("💬 Gemini RAG Response:\n", response)

# You can also test with other queries
print("\n" + "="*50)
query2 = "What are some good sources of protein with low calories?"
response2 = rag_query(query2)
print("\n🧠 Query:", query2)
print("💬 Gemini RAG Response:\n", response2)