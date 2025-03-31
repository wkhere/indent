build:
	go test -cover
	go build ./cmd/indent

install: build
	go install ./cmd/indent

cover:
	go test -coverprofile=cov
	go tool cover   -html=cov

sel=.
cnt=6
bench:
	go test -bench=$(sel) -count=$(cnt) -benchmem

.PHONY: build install cover bench
