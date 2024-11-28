# Commit Log Service

## Overview
The **Commit Log Service** is a distributed service built with Go, designed to replicate and demonstrate core functionalities of popular commit log systems like Apache Kafka or Pulsar. It serves as a learning tool to solidify concepts around building distributed systems and implementing core log-based message queues.

This project showcases a minimal, performant, and scalable commit log service to store and retrieve records in an append-only log structure, perfect for understanding the backbone of event-driven architectures.

---

## Features
- **Append-Only Log**: Records are written sequentially for efficient storage and retrieval.
- **Produce/Consume API**: Supports writing (producing) and reading (consuming) records via HTTP.
- **Distributed System Design**: Built with scalability and fault tolerance in mind.
- **Protobuf Integration**: Structured and efficient data serialization using Protocol Buffers.
- **Go-Powered**: Leverages Go's performance and concurrency model for robust service delivery.

---

## Getting Started

### Prerequisites
- **Go**: Ensure you have [Go](https://go.dev/) installed (v1.18+ recommended).
- **Protobuf**: Install the Protocol Buffers compiler `protoc`. Refer to the [official documentation](https://developers.google.com/protocol-buffers) for installation steps.

---

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/tunedev/commit-log-service.git
   cd commit-log-service
   ```

2. Build the service:
   ```bash
   go build -o commitlog
   ```

3. Start the service:
   ```bash
   ./commitlog
   ```

By default, the service listens on `http://localhost:8080`.

---

### Usage

#### Producing Records
Send a record to the log using the `/produce` endpoint:

```bash
curl -X POST http://localhost:8080/produce \
-H "Content-Type: application/json" \
-d '{"record": {"value": "hello world"}}'
```

#### Consuming Records
Retrieve a record by its offset using the `/consume` endpoint:

```bash
curl -X GET http://localhost:8080/consume \
-H "Content-Type: application/json" \
-d '{"offset": 0}'
```

---

## Code Structure

### Main Service
The service is initialized in the `main.go` file:
```go
package main

import (
	"log"

	"github.com/tunedev/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
```

### Protocol Buffers
The API interactions are defined using Protocol Buffers for efficient serialization. The `log.proto` file describes request and response messages:
```proto
syntax = "proto3";

package log.v1;

message ProduceRequest {
  Record record = 1;
}

message ProduceResponse {
  uint64 offset = 1;
}

message ConsumeRequest {
  uint64 offset = 1;
}

message ConsumeResponse {
  Record record = 2;
}

message Record {
  bytes value = 1;
  uint64 offset = 2;
}
```

---

## Contributions
We welcome contributions! Feel free to:
1. Fork the repository.
2. Create a new branch.
3. Submit a pull request.

---

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Author
Built and maintained by [Tunde Dev](https://github.com/tunedev).

Happy coding! 🚀
