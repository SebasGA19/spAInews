FROM node:alpine
# Setup user
RUN addgroup -g 5001 -S frontend
RUN adduser -h /apt/frontend -s /sbin/nologin -G frontend -S -u 5001 frontend
# Build the app
COPY frontend /apt/frontend/app
COPY configs/frontend/frontend.env /apt/frontend/app/.env
RUN chmod -R 777 /apt/frontend/
USER frontend:frontend
WORKDIR /apt/frontend/app
RUN npm install
RUN npm run build
CMD node build/index.js

