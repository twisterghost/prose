default:
	go build -o prose main.go

install:
	install prose /usr/local/bin
