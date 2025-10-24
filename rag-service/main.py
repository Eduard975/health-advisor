from fastapi import FastAPI
from pydantic import BaseModel
from rag import send_query 

app = FastAPI(title="RAG API")

class QueryRequest(BaseModel):
    query: str

class QueryResponse(BaseModel):
    answer: str

@app.get("/test", response_model=QueryResponse)
async def handle_query():
    return QueryResponse(answer="RAG service is up and running")

@app.post("/query", response_model=QueryResponse)
async def handle_query(request: QueryRequest):
    answer = send_query(request.query)
    return QueryResponse(answer = answer)
