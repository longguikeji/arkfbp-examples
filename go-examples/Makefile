override BIN_DIR      := bin
override APP_NAME     := app
override CURRENT_DIR  := $(shell pwd)
override SYSTEM_ARCH  := $(shell uname)
override TARGET := $(BIN_DIR)/$(APP_NAME)

all: $(APP_NAME)

$(APP_NAME): main.go cmd_root.go cmd_serve.go cmd_version.go cmd_run.go route.go
	@echo "Building $@ ..."
	@env \
		go build -o $(BIN_DIR)/$@ $^
	@echo "Build $@ finished!\n"

clean:
	@if [ -f ${TARGET} ] ; then rm ${TARGET} ; fi

.PHONY: clean
