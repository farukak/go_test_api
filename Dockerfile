FROM golang:1.15 AS build

WORKDIR /go_test_api

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

FROM alpine:latest

WORKDIR /app


COPY --from=build /go_test_api/app .
COPY --from=build /go_test_api/.env .


EXPOSE 3000

CMD ["./app"]