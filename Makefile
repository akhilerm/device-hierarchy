all: clean build image

clean:
	rm topo

build:
	@go build -o topo main.go

image: build
	@sudo docker build -t device-topology:ci .
