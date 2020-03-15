build:
	go build .

build_release:
	go build -ldflags "-s -w"

clean:
	rm quotes