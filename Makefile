file = cmd/hashdir.go

default:
	GOOS=linux GOARCH=amd64 go build -o build/hash_linux ${file}
	GOOS=windows GOARCH=amd64 go build -o build/hash_win ${file}
	GOOS=darwin GOARCH=amd64 go build -o build/hash_mac ${file}

