build:
	@go build -o ./bin/blockchainGo

run: build
	@./bin/blockchainGo

test:
	go test -v ./...