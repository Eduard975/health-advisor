from fastapi import FastAPI
from pydantic import BaseModel
from typing import List, Optional
from rag import send_query

app = FastAPI(title="RAG API")

class ChatHistoryItem(BaseModel):
    sender: str
    text: str
    timestamp: Optional[str] = None  

class QueryRequest(BaseModel):
    query: str
    history: List[ChatHistoryItem] = []  

class QueryResponse(BaseModel):
    answer: str

@app.get("/test", response_model=QueryResponse)
async def handle_test():
    return QueryResponse(answer="RAG service is up and running")

@app.post("/query", response_model=QueryResponse)
async def handle_query(request: QueryRequest):
    answer = send_query(request.query, request.history)
    return QueryResponse(answer=answer)
