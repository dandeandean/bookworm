FROM golang:1.25

WORKDIR /app
COPY .. /app
RUN go mod download && \
				CGO_ENABLED=0 GOOS=linux go build -o bookworm

CMD ["/app/bookworm"]
