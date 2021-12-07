build:
	go build ./...

vet:
	go vet ./...

ci: vet build
	@echo "validation succeeded"