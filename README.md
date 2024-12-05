
# Ethereum Tx Parser

This project implements an Ethereum blockchain parser that allows querying transactions for subscribed addresses.

## Features

- Fetch the current Ethereum block number using JSON-RPC.
- Subscribe to Ethereum addresses to monitor their transactions.
- Retrieve inbound and outbound transactions for subscribed addresses.

## Project Structure

The project is organized into a clean, layered architecture to ensure scalability, maintainability, and testability:

```plaintext
tx-parser/
├── main.go                # Entry point of the application
├── config/                # Configuration-related code
│   └── config.go
├── domain/                # Core business models and interfaces
│   ├── models.go          # Data structures and domain entities
│   ├── parser.go          # Interface for the Parser service
│   └── repository.go      # Interface for Repository abstraction
├── service/               # Business logic and interaction with repositories
│   ├── parser_service.go  # Implementation of the Parser service
│   └── parser_service_test.go # Unit tests for ParserService
├── repository/            # Data persistence layer
│   ├── memory_repository.go   # In-memory implementation of Repository
│   └── memory_repository_test.go # Unit tests for MemoryRepository
├── api/                   # HTTP API layer
│   ├── http_handler.go    # API handlers for HTTP endpoints
│   ├── http_handler_test.go # Unit tests for API handlers
│   └── models.go          # Models used in HTTP responses
├── utils/                 # Helper utilities (e.g., logging, JSON-RPC)
│   ├── converter.go       # Hex string to integer or related conversions
│   ├── converter_test.go  # Unit tests for Converter utility
│   ├── jsonrpc.go         # JSON-RPC communication helper
│   └── logger.go          # Logger setup and utilities
├── middleware/            # Middleware for HTTP server
│   ├── log_request.go     # Logs every incoming HTTP request
│   └── recovery.go        # Panic recovery middleware
├── mocks/                 # Auto-generated mocks for testing
│   ├── Parser.go          # Mock for Parser interface
│   └── Repository.go      # Mock for Repository interface
├── scripts/               # Helper scripts for testing or deployment
│   └── tests.py           # Python script for testing APIs
├── .env                   # Environment variables file
├── go.mod                 # Go module definition
├── go.sum                 # Go module dependencies
├── makefile               # Commands to build, run, and test the app
├── README.md              # Project documentation
└── app.log                # Log file generated during runtime

```

## Endpoints

### 1. Get Current Block

Fetches the latest block number from the Ethereum blockchain.

- **Endpoint:** `/currentBlock`
- **Method:** `GET`
- **Response:**
  ```json
  {
    "currentBlock": 123456
  }
  ```

### 2. Subscribe to Address

Subscribes an Ethereum address for monitoring transactions.

- **Endpoint:** `/subscribe`
- **Method:** `GET`
- **Query Parameters:**
  - `address`: The Ethereum address to subscribe.
- **Response:**
  - `200 OK` on success.
  - `400 Bad Request` if the address is missing.

### 3. Get Transactions

Retrieves transactions for a subscribed Ethereum address.

- **Endpoint:** `/transactions`
- **Method:** `GET`
- **Query Parameters:**
  - `address`: The Ethereum address.
- **Response:**
  ```json
  [
    {
      "from": "0xSenderAddress",
      "to": "0xReceiverAddress",
      "value": "100",
      "hash": "0xTransactionHash"
    }
  ]
  ```

## Prerequisites

1. Install [Go](https://golang.org/dl/) (version 1.16 or higher).

## Running the Project

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd tx-parser
   ```

2. Run the program:
   ```bash
   go run main.go
   ```

## Notes

- The project uses in-memory storage, but it can be extended for other storage systems.
- Ethereum JSON-RPC is used for blockchain interaction.
