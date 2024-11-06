# Weather Service

A gRPC-based weather service providing weather information using Protocol Buffers and Connect RPC.

## Features

- Temperature readings
- Weather descriptions
- Location data
- Humidity levels
- Wind speed information

## Project Structure

├── proto/ # Protocol Buffer definitions
├── gen/ # Generated code
├── handler/ # Service implementation
├── interceptors/ # Logging middleware
└── main.go # Entry point

## Getting Started

### Prerequisites

- Go 1.16+
- buf CLI

### Available Commands

```bash
# Generate code from protos
make generate

# Clean up the generated code
make clean

# Build the project
make build

# Run the server
make run
```

### Running the Service

Start the server:

`make run`
