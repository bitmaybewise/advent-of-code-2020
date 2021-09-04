default: test
test:
	go test -race ./...
test1:
	go test ./... -run=$(filter)
run:
	go run src/day$(day)/main.go
