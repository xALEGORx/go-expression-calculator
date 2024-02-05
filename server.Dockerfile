FROM node:16-alpine
WORKDIR /app
COPY frontend/package*.json ./
RUN npm install
COPY frontend ./
RUN npm run build

FROM nginx:alpine
COPY --from=0 /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

FROM golang:1.20 AS orchestrator
WORKDIR /app
COPY go.* ./
COPY internal ./internal
RUN go mod download
COPY cmd/orchestrator ./
RUN CGO_ENABLED=0 GOOS=linux go build -o orchestrator .

FROM scratch
COPY --from=orchestrator /app /orchestrator
CMD ["/orchestrator"]