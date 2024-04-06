FROM node:20-slim as web

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install

COPY . ./

ENTRYPOINT ["/app/node_modules/.bin/vite"]
