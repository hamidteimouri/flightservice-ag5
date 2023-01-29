FROM golang:1.19-alpine AS build

WORKDIR /app/ht

COPY /src/go.mod .
COPY /src/go.sum .

RUN go mod download

COPY . .
COPY /src/cmd/.env.example ./.env
WORKDIR /app/ht/src/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./app

# gRPC port
EXPOSE 50051

## Run the binary
ENTRYPOINT ["./app"]
