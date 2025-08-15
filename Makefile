run:
	go run gia.go

build:
	go build -o gia gia.go

test:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

clean-build:
	rm -f gia