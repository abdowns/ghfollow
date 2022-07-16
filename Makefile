build:
	go build -o ghfollow src/*.go

test: build
	./ghfollow