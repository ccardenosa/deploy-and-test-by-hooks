GOCMD=go
MOD_NAME=github.com/ccardenosa/front-back-app

SUBDIRS := backend 
# SUBDIRS := $(wildcard */.)


all: run

$(SUBDIRS):
	$(MAKE) -C $@ $<

setup:
	test -f go.mod || $(GOCMD) mod init $(MOD_NAME)
	$(foreach var,$(SUBDIRS), \
			$(GOCMD) mod edit -replace $(MOD_NAME)/$(var)=./$(var); \
	 )
	$(GOCMD) mod tidy

linters-install:
	@golangci-lint --version >/dev/null 2>&1 || { \
		echo "installing linting tools..."; \
		curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s v1.41.1; \
	}

lint: linters-install
	bin/golangci-lint run

run: $(SUBDIRS) setup
	$(GOCMD) run .

test: $(SUBDIRS) setup
	$(GOCMD) test -v ./...

coverage: $(SUBDIRS) setup
	$(GOCMD) test -cover -race ./...

bench:
	$(GOCMD) test -run=NONE -bench=. -benchmem ./...

clean:
	rm -f go.mod go.sum
	$(foreach var,$(SUBDIRS), \
			make -f $(var)/Makefile clean; \
	 )

.PHONY: all run test coverage lint linters-install $(SUBDIRS) clean
