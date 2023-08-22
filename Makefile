
# Go-GPT CLI installation
go-gpt:
	@echo "Installing Go-GPT...."
	@go get .

build:
	@go build -o bin/go-gpt cmd/go_gpt/go_gpt.go

# Generate build for cross-compilation
go-build-win:
	@GOOS=windows GOARCH=amd64 go build -o bin/go-gpt cmd/go_gpt/go_gpt.go

go-build-linux:
	@GOOS=linux GOARCH=amd64 go build -o bin/linux/go-gpt cmd/go_gpt/go_gpt.go

go-build-mac:
	@GOOS=darwin GOARCH=amd64 go build -o bin/macos/go-gpt cmd/go_gpt/go_gpt.go

tidy: ## add missing and remove unused modules
	 go mod tidy

install-deps: | install-gofumpt ## install some project dependencies

install-gofumpt:
	# install gofumpt
	go install mvdan.cc/gofumpt@latest