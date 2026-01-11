BUILDER ?= podman

all:
	go build

install:
	go install

build:
	go build . -o bookworm

test:
	go test ./... -v

docker:
	$(BUILDER) build -f ./build/Dockerfile . -t bookworm

nix:
	nix-build ./build/bookworm.nix
