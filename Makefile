run:
	go run main.go

build: 
	go build -o app .

test: 
	go test ./..

clean:
	rm -f app