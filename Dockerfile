FROM golang:1.19 AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN go mod verify

RUN GOOS=linux go build -a -o /api bin/cmd/main.go

FROM alpine:latest AS prod

WORKDIR /app

COPY --from=build /app/api .
EXPOSE 3030
ENTRYPOINT [ "./api" ]

