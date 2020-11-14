build:
	go build -o bin/main main.go

run:
	go run main.go 2 5 10

test:
	go test -v

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=386 go build -o bin/mt-trainer-linux-386 main.go
	GOOS=linux GOARCH=arm go build -o bin/mt-trainer-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/mt-trainer-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/mt-trainer-freebsd-386 main.go