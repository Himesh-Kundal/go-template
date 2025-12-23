FROM golang:1.24-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o api ./cmd/api

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/api .
EXPOSE 8080
CMD ["./api"]
