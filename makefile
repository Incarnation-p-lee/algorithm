SRC_DIR  :=./src/
PKG_DIR  := ./pkg/
PACKAGES :=number

.PHONY: install test update help

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
	@export GOPATH=`pwd`
	@cd $(SRC_DIR) && glide up @cd -

install:
	@export GOPATH=`pwd`
	@rm -rf $(PKG_DIR)
	go install -buildmode=shared -linkshared $(PACKAGES)

test:
	export GOPATH=`pwd`

