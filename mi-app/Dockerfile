FROM node:latest

WORKDIR /app

COPY package.json . 
COPY package-lock.json .

RUN npm install --force 
RUN npm install -g vite

COPY . . 

EXPOSE 5173

CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0", "--port", "5173"]
