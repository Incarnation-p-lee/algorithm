SRC_DIR   :=./src/
PKG_DIR   := ./pkg/
PACKAGES  :=number assert
PROFILES  :=$(addsuffix .out, $(addsuffix .cover, $(PACKAGES)))
SRC_FILES :=$(shell find $(SRC_DIR) -name *.go | grep -v vendor)

.PHONY: install test update help $(PACKAGES)

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
	@export GOPATH=`pwd` && cd $(SRC_DIR) && glide up cd -

install:
	@rm -rf $(PKG_DIR)
	@export GOPATH=`pwd` \
        && gofmt -w $(SRC_FILES) \
        && go install -buildmode=shared -linkshared $(PACKAGES)

test:$(PROFILES)
	@export GOPATH=`pwd` && gofmt -w $(SRC_FILES)

$(PROFILES):%.cover.out:%
	@export GOPATH=`pwd` && go test -v -coverprofile $@ $<

