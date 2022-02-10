
APP="Project2"
ALL_PACKAGES=$(shell go list ./... | grep -v "docs")
APP_EXECUTABLE="./out/$(APP)"

setup:
	go get -d golang.org/x/lint/golint
	go get -u -v github.com/go-swagger/go-swagger/cmd/swagger
	brew install golang-migrate
	go install golang.org/x/tools/cmd/goimports@latest

lean:
	rm -rf out/
	rm -f *.out
	go clean -testcache

fmt:
	@echo "running fmt..."
	@go fmt ./...

vet:
	@echo "running vet..."
	@go vet ./...

lint:
	@echo "running lint..."
	@for p in $(ALL_PACKAGES); do \
  		golint $$p | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; } \
  	done

compile:
	@mkdir -p out/
	@go build -o $(APP_EXECUTABLE)

build: fmt lint vet compile

test:
	@go test $(ALL_PACKAGES)

test-cover:
	@mkdir -p out/
	@go test $(ALL_PACKAGES) -coverprofile=coverage.out

test-cover-report: test-cover
	@echo 'Total coverage: '`go tool cover -func coverage.out | tail -1 | awk '{print $$3}'`

test-cover-html: test-cover
	@go tool cover -html=coverage.out

createDB:
	psql -h localhost -U postgres -w -c "create database test;"

dropDB:
	psql -h localhost -U postgres -w -c "drop database test;"

migrateUP:
	migrate -path db/migrations -database "postgres://postgres:0712@localhost:5432/test?sslmode=disable" -verbose up

migrateDOWN:
	migrate -path db/migrations -database "postgres://postgres:0712@localhost:5432/test?sslmode=disable" -verbose down
	
.PHOBY: createDB dropDB migrateUP migrateDOWN