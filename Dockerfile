FROM golang:1.23.4 AS builder
WORKDIR /app
COPY go.mod go.sum .env ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/main /main
COPY .env .env
EXPOSE 5001
CMD ["/main"]
