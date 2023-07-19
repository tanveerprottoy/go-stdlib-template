build:
	go build -o bin/app cmd/template/main.go

run:
	go run cmd/template/main.go

test-run:
	go test -run

# package wise test
test-user:
	go test ./internal/app/template/user

# specific tese function with package path
# go test -run TestMultiply ./