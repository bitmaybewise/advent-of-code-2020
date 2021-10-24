default: test
test:
	go test -race ./...
test1:
	go test ./... -run=$(filter) -timeout=1s
run:
	go run day$(day)/main.go
