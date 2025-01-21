run:
	go run .

build:
	go build -o build/gterm -ldflags="-s -w"

todo:
	@grep -rn -r --exclude=\Makefile '// TODO:'

test:
	@go test -failfast $(go list ./... | grep -v ./internal/ui) -count=1

coverage:
	@go test $(go list ./... | grep -v ./internal/ui)  -coverpkg=$(go list ./... | grep -v ./internal/ui). -coverprofile ./coverage.out && clear && go tool cover -func ./coverage.out