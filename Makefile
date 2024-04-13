run-comment:
	go run cmd/comment/main.go

test-coverage:
	mkdir -p coverage
	go test -race -short -v -coverprofile coverage/cover.out ./...
	go tool cover -html=coverage/cover.out

gen-swag:
	swag init -d ./cmd/comment,./handler/commenthandler -o ./cmd/comment/doc --pd