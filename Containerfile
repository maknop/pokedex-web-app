FROM golang:1.19

WORKDIR /

COPY go.mod go.sum /
RUN go mod download

COPY . /

RUN go build -o /pokedex-web-app

EXPOSE 8080

CMD ["/pokedex-web-app"]