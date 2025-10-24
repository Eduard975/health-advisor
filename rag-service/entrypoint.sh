#!/bin/bash
set -e

# Ensure vectors folder exists
mkdir -p /app/data/vectors

# Generate vectorstore if empty
if [ -z "$(ls -A /app/data/vectors)" ]; then
    echo "Vectors folder empty — generating..."
    python createVectorstore.py
else
    echo "Vectors folder already exists — skipping generation"
fi

# Start FastAPI
uvicorn main:app --host 0.0.0.0 --port 8000 --reload