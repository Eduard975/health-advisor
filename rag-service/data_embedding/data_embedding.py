from langchain_text_splitters import RecursiveCharacterTextSplitter
from langchain_openai import OpenAIEmbeddings
from langchain_community.vectorstores import FAISS

import pandas as pd

df = pd.read_csv("data/food/food_nutrition_text.csv")
texts = df["text"].tolist()

splitter = RecursiveCharacterTextSplitter(chunk_size=1000, chunk_overlap=100)
docs = splitter.create_documents(texts)

embeddings = OpenAIEmbeddings(model="text-embedding-3-large")
db = FAISS.from_documents(docs, embeddings)
db.save_local("data/food/food_vectorstore")
