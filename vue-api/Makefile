SHELL=cmd.exe
DSN=host=192.168.88.144 port=5437 user=postgres password=password dbname=vueapi sslmode=disable timezone=UTC connect_timeout=5
BINARY_NAME=vueapi.exe
ENV=development

## build: builds all binaries
build:
	@go build -o ${BINARY_NAME} ./cmd/api
	@echo back end built!

run: build
	@set DSN=${DSN}&& set ENV=${ENV}&& start /B .\${BINARY_NAME} &
	@echo back end started!

clean:
	@echo Cleaning...
	@DEL ${BINARY_NAME}
	@go clean
	@echo Cleaned!

start: run

stop:
	@taskkill /IM ${BINARY_NAME} /F
	@echo back end stopped! 

restart: stop start

