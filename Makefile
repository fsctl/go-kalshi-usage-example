.PHONY: clean all

GOFLAGS ?=

all:
	$(MAKE) -C cmd/kot
	$(MAKE) -C cmd/list-markets

clean:
	@rm -rf kot
	@rm -rf list-markets

