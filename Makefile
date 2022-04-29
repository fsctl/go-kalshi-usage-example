.PHONY: clean all

GOFLAGS ?=

all:
	$(MAKE) -C cmd/list-markets

clean:
	@rm -rf list-markets

