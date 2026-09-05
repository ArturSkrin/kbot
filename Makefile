REGISTRY   ?= ghcr.io
REPOSITORY ?= arturskrin/kbot
APP        ?= kbot
VERSION    ?= v1.0.0
TARGETOS   ?= linux
TARGETARCH ?= amd64
SHA        := $(shell git rev-parse --short HEAD)
TAG        := $(VERSION)-$(SHA)
IMAGE      := $(REGISTRY)/$(REPOSITORY):$(TAG)-$(TARGETOS)-$(TARGETARCH)

.PHONY: format lint test get build image push clean

format:
	gofmt -s -w ./

lint:
	go vet ./...

test:
	go test ./...

get:
	go mod download

build:
	CGO_ENABLED=0 GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) go build -o $(APP) .

image:
	docker build --platform $(TARGETOS)/$(TARGETARCH) -t $(IMAGE) .

push:
	docker push $(IMAGE)

clean:
	rm -f $(APP)
	@echo cleaned