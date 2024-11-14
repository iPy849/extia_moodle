#Change these variables as necessary.
MAIN_PACKAGE_PATH := .
BINARY_NAME := extia
COMPILATION_PATH := ./bin
TEST_PATH := ${COMPILATION_PATH}/test_reports

update:
	@echo "Dependency update"
	go mod tidy -v

test:
	@echo "Running tests"
	@mkdir -p bin/test_reports
	go test -v -buildvcs -coverprofile=${TEST_PATH}/coverage.out ./...


test-report: test
	@echo "Running tests report"
	go test -v -buildvcs -coverprofile=${TEST_PATH}/coverage.out ./...
	go tool cover -html=${TEST_PATH}/coverage.out

build:
	@echo "Build project"
	@mkdir -p bin
	go build -o=${COMPILATION_PATH}/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

build-pi:
	@echo "Build project for Raspberry"
	@mkdir -p bin
	env CGO_ENABLED=1 GOARCH=arm64 GOOS=linux go build -o=${COMPILATION_PATH}/${BINARY_NAME}.pi ${MAIN_PACKAGE_PATH}

run: clear build
	@echo "Running project"
	${COMPILATION_PATH}/${BINARY_NAME}

clear:
	@echo "Removing builds and tests"
	@rm -rf ./bin


all: clear update test run
