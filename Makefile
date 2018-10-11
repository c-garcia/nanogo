.PHONY: test

dist/c1: $(wildcard ./cmd/c1/*.go) dist
	@go build -o dist/c1 ./cmd/c1

dist:
	@mkdir dist	
test:
	@go test

clean:
	@rm -rf dist
