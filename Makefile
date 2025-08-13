run:
	go run gia.go

build:
	go build -o gia gia.go

install: build
	sudo cp gia /usr/local/bin/gia

uninstall:
	sudo rm /usr/local/bin/gia