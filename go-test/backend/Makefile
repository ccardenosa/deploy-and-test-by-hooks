GOCMD=go
CUR_DIR=$(shell pwd)
MOD_NAME=github.com/ccardenosa/front-back-app/$(shell basename $(CUR_DIR))

frontend:
	test -f go.mod || $(GOCMD) mod init $(MOD_NAME)
	$(GOCMD) mod tidy

clean:
	rm -f go.mod go.sum
