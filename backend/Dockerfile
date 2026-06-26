FROM golang:1.26-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /server .

FROM alpine:3.22

WORKDIR /app

COPY --from=build /server /app/server

EXPOSE 8080

CMD ["/app/server"]
