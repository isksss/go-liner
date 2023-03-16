.PHONY: build tidy

tidy:
	go mod tidy
	cd cmd && go mod tidy

build: tidy
	cd cmd/liner && go build
	mkdir -p dist
	mv ./cmd/liner/liner ./dist/liner