from langchain.chains import RetrievalQA
from langchain.chat_models import ChatOpenAI

retriever = db.as_retriever(search_kwargs={"k": 3})
llm = ChatOpenAI(model="gpt-5", temperature=0)
rag = RetrievalQA.from_chain_type(llm=llm, retriever=retriever)

query = "Which foods are highest in Vitamin C and lowest in fat?"
print(rag.run(query))