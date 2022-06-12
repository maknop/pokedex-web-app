fmt:
	go fmt ./...
	
vet:
	go vet ./...

install: fmt vet
	go mod download

run:
	go run *.go
