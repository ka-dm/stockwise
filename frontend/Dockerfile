FROM node:20-alpine3.19
WORKDIR /app/stockwise
COPY stockwise/package*.json ./
RUN npm install
COPY stockwise .
EXPOSE 5173
CMD ["npm", "run", "dev", "--", "--host"]
