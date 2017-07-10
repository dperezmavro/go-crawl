.PHONY : all

all: main
	GOARCH=amd64 GOOS=darwin go build -ldflags="-s -w" -o main *.go