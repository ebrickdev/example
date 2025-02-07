# Ebrick Sample

## Getting Started

Follow these steps to set up and run the project:

### Initialize the Project

```bash
# Initialize the project and generate protobuf files
make init proto
```

### Update and Start Docker Containers

```bash
# Update and start the Docker containers in detached mode
docker compose update -d
```

### Run the Application

```bash
# Run the main Go application
go run main.go
```

## Guidelines

1. Ensure you have Docker and Go installed on your machine.
2. Run `make init proto` to initialize the project and generate necessary protobuf files.
3. Use `docker compose update -d` to update and start the Docker containers.
4. Execute `go run main.go` to start the main Go application.
5. Refer to the project's documentation for more detailed instructions and troubleshooting.

