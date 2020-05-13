default:
	go build -o bin/prose main.go

clean:
	rm -rf bin

install:
	install bin/prose /usr/local/bin
