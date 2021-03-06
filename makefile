SRC_DIR   :=src/
PKG_DIR   :=pkg/
PACKAGES  :=number assert map-reduce/task
PROFILES  :=$(addsuffix .cover.out, $(PACKAGES))
SRC_FILES :=$(shell find $(SRC_DIR) -name *.go | grep -v vendor)
GOPATH    :=$(shell pwd)

.PHONY: install test update help $(PACKAGES)

install:
	@rm -rf $(PKG_DIR)
	@export GOPATH=$(GOPATH) \
        && gofmt -w $(SRC_FILES) \
        && go install -buildmode=shared -linkshared $(PACKAGES)

help:
	@echo Usage
	@echo
	@echo "	make [command]"
	@echo
	@echo "The commands are:"
	@echo
	@echo "	help:    print the help message."
	@echo "	update:  update dependency packages by glide."
	@echo "	install: build dynamic and static library and install into $(PKG_DIR)."
	@echo "	test:    run test code."
	@echo

update:
	@export GOPATH=$(GOPATH) && cd $(SRC_DIR) && glide up cd -

test:$(PROFILES)
	@export GOPATH=$(GOPATH) && gofmt -w $(SRC_FILES)
	@cat *.cover.out > cover.out
	@rm *.cover.out

$(PROFILES):%.cover.out:%
	@export GOPATH=`pwd` \
        && go test -v -coverprofile $$$$.cover.out -covermode=atomic $<

