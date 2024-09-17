# Build the application
all: build

build:
	@echo "Building $(APP)..."
	@if [ "$(APP)" = "project" ]; then \
		go build -o target/project-main cmd/project/*.go; \
	elif [ "$(APP)" = "grammar" ]; then \
		go build -o target/grammar-main cmd/grammar/*.go; \
	else \
		echo "Error: Unknown application. Use 'make run APP=<balancer, shuntingyard, ast, afn, project>'"; \
		exit 1; \
	fi

# Run the application
run: build
	@echo "Running $(APP)..."
	@if [ "$(APP)" = "project" ]; then \
		./target/project-main; \
	elif [ "$(APP)" = "grammar" ]; then \
		./target/grammar-main; \
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