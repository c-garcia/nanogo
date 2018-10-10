.PHONY: test

dist/c1: dist
	@go build -o dist/c1 ./cmd/c1

dist:
	@mkdir dist	
test:
	@go test

clean:
	@rm -rf dist
