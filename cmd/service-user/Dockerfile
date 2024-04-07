FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
COPY . .

RUN go mod tidy && go build -o service-user .
RUN go install github.com/cosmtrek/air@v1.51.0

EXPOSE 3000

CMD ["air", "-c", ".air.toml"]