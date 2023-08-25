FROM golang:1.19 AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN go mod verify

RUN GOOS=linux go build -a -o /api bin/cmd/main.go
EXPOSE 3030
ENTRYPOINT [ "/api" ]
