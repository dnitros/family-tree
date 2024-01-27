BIN_DIR = /usr/local/bin

help:          ## Show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

install:       ## Install Target
	@go build -o ${BIN_DIR}/family-tree

uninstall:     ## Uninstall Target
	@rm -f ${BIN_DIR}/family-tree