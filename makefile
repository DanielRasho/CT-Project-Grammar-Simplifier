# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building $(APP)..."
	@if [ "$(APP)" = "balancer" ]; then \
		go build -o target/balancer-main cmd/balancer/*.go; \
	elif [ "$(APP)" = "shuntingyard" ]; then \
		go build -o target/shuntingyard-main cmd/ShuntingYard/*.go; \
	elif [ "$(APP)" = "ast" ]; then \
		go build -o target/ast-main cmd/abstract_syntax_tree/*.go; \
	elif [ "$(APP)" = "afn" ]; then \
		go build -o target/thompson-main cmd/nondeterministic_finite_automaton/*.go; \
	elif [ "$(APP)" = "project" ]; then \
		go build -o target/project-main cmd/project/*.go; \
	else \
		echo "Error: Unknown application. Use 'make run APP=<balancer, shuntingyard, ast, afn, project>'"; \
		exit 1; \
	fi

# Run the application
run: build
	@echo "Running $(APP)..."
	@if [ "$(APP)" = "balancer" ]; then \
		./target/balancer-main; \
	elif [ "$(APP)" = "shuntingyard" ]; then \
		./target/shuntingyard-main; \
	elif [ "$(APP)" = "ast" ]; then \
		./target/ast-main; \
	elif [ "$(APP)" = "afn" ]; then \
		./target/thompson-main; \
	elif [ "$(APP)" = "project" ]; then \
		./target/project-main; \
	else \
		echo "Error: Unknown application. Use 'make run APP=<balancer, shuntingyard, ast, afn, project>'"; \
		exit 1; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./tests -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f target/*