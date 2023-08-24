FROM golang:1.19 AS build

WORKDIR /app
COPY . .

COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

RUN go build -a -o api bin/cmd/main.go

FROM alpine:latest AS prod

WORKDIR /app

COPY --from=build /app/api .

ENTRYPOINT [ "./api" ]
