build:
	go build -o bin/app cmd/template/main.go

run:
	go run cmd/template/main.go

test-run:
	go test -v -run

# package wise test
test-user:
	go test -v -cover ./internal/app/template/module/user

# specific tese function with package path
# go test -run TestMultiply ./

# go test -v <package> -run <TestFunction>
# go test -v -cover --short -race  ./... -run ^TestError*