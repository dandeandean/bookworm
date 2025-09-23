all:
	go build

install:
	go install

build:
	go build . -o bookworm

test:
	go test ./... -v

docker:
	docker build -f ./build/Dockerfile . -t bookworm

nix:
	nix-bulid build/
