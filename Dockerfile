FROM golang:1.19.3

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o pc/gRPC-Go-Gorm

EXPOSE 8000

CMD ["pc/gRPC-Go-Gorm"]