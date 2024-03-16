GO=go

run:
	${GO} run ./cmd

build:
	${GO} build -o ./cmd/ciaoExecutableFile ./cmd/main.go

buildrun:
	./cmd/ciaoExecutableFile

swaggo:
	swag init -g ./internal/infrastructure/api/server.go

swaggoformat:
	swag fmt	

test:
	${GO} test -v ./...
		