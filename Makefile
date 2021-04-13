# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOBUILDPLUGIN=$(GOBUILD) -buildmode=plugin
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOGENERATE=$(GOCMD) generate
GOLIST=$(GOCMD) list


# Sources parameters
SOURCE_ENTRYPOINT=./cmd/main.go
PLUGIN_SOURCE_ENTRYPOINT=./puzzles/2020/01/main.go

# Binary parameters
BINARY_NAME=solver
BINARY_DESTINATION=./bin
BINARY_PATH=$(BINARY_DESTINATION)/$(BINARY_NAME)

ifeq (run,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

# Tagets
help:
	@echo "Compilation, build, run"
	@echo "build               : Build the app"
	@echo "run 01              : Run the app with the day of your choice"
	@echo "test                : Launch all units tests"
	@echo "test_cover          : Launch all units tests with the cover of each package"
	@echo "clean               : Remove temporary files"

build:
		$(GOBUILD) -o $(BINARY_PATH) -v $(SOURCE_ENTRYPOINT)
test:
		$(GOTEST) -v ./...
test_cover:
		$(GOTEST) -v ./... -coverprofile=coverage.txt -covermode=atomic
run:
	[ -n "$(RUN_ARGS)" ] || exit 1
	@echo "Running Advent of Code for December $(RUN_ARGS)"
	$(GORUN) $(SOURCE_ENTRYPOINT) -day $(RUN_ARGS)
clean:
		$(GOCLEAN) $(SOURCE_ENTRYPOINT)
		rm -f $(BINARY_PATH)
