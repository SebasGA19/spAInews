# Build
FROM node:alpine as builder
# Build the app
COPY frontend /frontend
COPY configs/frontend/frontend.env /frontend/.env

WORKDIR /frontend
RUN npm install
RUN npm run build

# Serve
FROM nginx:alpine as server

WORKDIR /usr/share/nginx/html
RUN rm -rf ./*
COPY --from=builder /frontend/build .
ENTRYPOINT ["nginx", "-g", "daemon off;"]
