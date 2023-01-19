test:
	go test ./... -short -coverprofile=cov.out `go list ./... | grep -v vendor/`
