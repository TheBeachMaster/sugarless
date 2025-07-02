.PHONY: clean build

APP_NAME = sugarless
BUILD_DIR = $(PWD)/bin

clean:
	rm -rf $(BUILD_DIR)

build: main.go
	go build -ldflags="-w -s -checklinkname=0" -o $(BUILD_DIR)/$(APP_NAME) main.go
