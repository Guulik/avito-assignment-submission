FROM golang:latest AS builder
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/banner

FROM alpine:latest

RUN apk --no-cache add ca-certificates postgresql-client redis

COPY --from=builder /app/app .

EXPOSE 4444

ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=5379
ENV POSTGRES_DB=bannerDb

CMD ["./app"]