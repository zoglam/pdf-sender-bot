FROM node:19-alpine3.17 as builder
WORKDIR /app
COPY package.json yarn.lock vite.config.ts ./
RUN yarn install
COPY . .
RUN yarn build

FROM node:19-alpine3.17
COPY --from=builder /app/dist /app/static