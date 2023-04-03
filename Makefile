
# Go-GPT CLI installation
go-gpt:
	@echo "Installing Go-GPT...."

# Generate build for cross-compilation
go-build-win:
	@GOOS=windows GOARCH=amd64 go build -o bin/go-gpt cmd/go_gpt.go

go-build-linux:
	@GOOS=linux GOARCH=amd64 go build -o bin/go-gpt cmd/go_gpt.go
