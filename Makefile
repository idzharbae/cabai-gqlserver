# go build command
build:
	@echo " >> building binaries"
	@go build -v -o gqlserver main.go

# go run command
run: build
	@./gqlserver

dep:
	@dep ensure -v