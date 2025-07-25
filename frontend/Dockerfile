FROM node:22-alpine3.19
WORKDIR /app/stockwise

RUN apk add --no-cache python3 make g++

COPY stockwise/package*.json ./
RUN npm install
COPY stockwise .
EXPOSE 5173
CMD ["npm", "run", "dev", "--", "--host"]
