from langchain_google_genai import GoogleGenerativeAIEmbeddings, ChatGoogleGenerativeAI
from langchain.chains import RetrievalQA

retriever = db.as_retriever(search_kwargs={"k": 3})

llm = ChatGoogleGenerativeAI(
    model="gemini-2.5-flash",
    temperature=0.2,
)

rag = RetrievalQA.from_chain_type(llm=llm, retriever=retriever)

query = "Which foods are highest in Vitamin C and lowest in fat?"
response = rag.run(query)

print("\n🧠 Query:", query)
print("💬 Gemini RAG Response:\n", response)