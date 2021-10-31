build:
	go build ./...

vet:
	go vet ./...

ci: vet build