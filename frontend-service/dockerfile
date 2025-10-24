# Use Node.js for dev
FROM node:20-alpine

WORKDIR /app

# Install dependencies
COPY package.json package-lock.json* ./
RUN npm ci

# Copy source code
COPY . .

# Expose port for Vite dev server
EXPOSE 5173

# Run Vite dev server
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0", "--port", "5173"]
