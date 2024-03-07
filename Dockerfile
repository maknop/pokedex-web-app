FROM golang:1.19

WORKDIR /

COPY go.mod go.sum /
RUN go mod download

COPY . /

# RUN go build -o /pokedex-web-app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build

EXPOSE 8080

CMD ["/pokedex-web-app"]
