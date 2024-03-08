FROM golang:1.19 as build

WORKDIR /

COPY go.mod go.sum /
RUN go mod download

COPY . /

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build

FROM debian:buster-slim

WORKDIR /
COPY --from=build / .

EXPOSE 8080

CMD ["/pokedex-web-app"]
