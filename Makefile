default: test

.PHONY: test
test:
	go test -race -v ./...

.PHONY: demo
demo:
	clear && /usr/bin/ls ./_examples | xargs -o -Ipkg sh -c 'echo && go run ./_examples/pkg'