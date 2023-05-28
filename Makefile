go:
	go vet
	go test -cover
	go install

cover:
	go test -coverprofile=cov
	go tool cover   -html=cov

bench:
	go test -bench=. -benchmem

.PHONY: go cover bench
