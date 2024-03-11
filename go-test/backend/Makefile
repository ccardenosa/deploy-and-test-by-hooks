GOCMD=go
CUR_DIR=$(shell pwd)
MOD_NAME=github.com/ccardenosa/deploy-and-test-by-hooks/go-test/$(shell basename $(CUR_DIR))

backend:
	test -f go.mod || $(GOCMD) mod init $(MOD_NAME)
	$(GOCMD) mod tidy

clean:
	rm -f go.mod go.sum
