build:
	go test -cover
	go build

install: build
	go install

cover:
	go test -coverprofile=cov
	go tool cover   -html=cov

bench:
	go test -bench=. -benchmem

.PHONY: build install cover bench
