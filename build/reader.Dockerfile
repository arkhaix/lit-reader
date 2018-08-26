# Build stage
FROM node:8.11 as builder
WORKDIR /usr/src/app

COPY web/app/package.json web/app/yarn.lock ./
RUN yarn

COPY web/app/ ./
RUN yarn build

# Result stage
FROM nginx:alpine
COPY --from=builder /usr/src/app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]