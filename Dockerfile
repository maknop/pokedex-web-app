FROM golang:1.19.2-alpine3.15

WORKDIR $GOPATH/src/github.com/maknop/pokedex-web-app

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

# Run the executable
CMD ["pokedex-web-app"]