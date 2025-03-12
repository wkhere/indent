build:
	go test -cover
	go build ./cmd/indent

install: build
	go install ./cmd/indent

cover:
	go test -coverprofile=cov
	go tool cover   -html=cov

bench:
	go test -bench=. -benchmem

.PHONY: build install cover bench
