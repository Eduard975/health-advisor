FROM python:3.11-slim

# Set working directory
WORKDIR /app

# Install system dependencies
RUN apt-get update && apt-get install -y \
    build-essential \
    && rm -rf /var/lib/apt/lists/*

# Copy requirements first for better caching
COPY requirements.txt .
COPY .env .
COPY entrypoint.sh .
COPY createVectorstore.py .

# Install Python dependencies
RUN pip install -r requirements.txt

# Set environment variables
ENV PYTHONUNBUFFERED=1

RUN chmod +x entrypoint.sh

# Expose port (adjust if your service uses a specific port)
EXPOSE 8000

ENTRYPOINT ["./entrypoint.sh"]