# IT Challenge: Proxy Server

This project is a reverse proxy server built in Golang. It forwards incoming HTTP requests to backend services, caches responses for improved performance, and allows SSL redirection. The project also includes a `/ping` endpoint for health checks.

## Features
- Reverse proxy for multiple backend services.
- Caching of backend responses with customizable expiration time.
- SSL/TLS support.
- Simple health check endpoint (`/ping`).
- Easy integration with Docker containers for backend services.

## Project Structure

- `cmd/main.go` — Main entry point to start the proxy server.
- `internal/server/` — Core server logic, including caching, proxy, and request handling.
- `internal/configs/` — Configuration management (via `config.yaml` file).
- `settings/config.yaml` — Configuration file for defining backend services.

## Getting Started

### Prerequisites

Ensure you have the following installed:

- [Go](https://golang.org/doc/install) (1.18 or later)
- [Docker](https://www.docker.com/get-started)
- [Make](https://www.gnu.org/software/make/)

### Installation

1. Clone the repository:

```bash
git clone https://github.com/Agoudjiliss/IT-challenge.git
cd IT-challenge
```

2. Install dependencies:

No additional dependencies are required beyond standard Go libraries and Docker.

### Configuration

1. Create a `config.yaml` file in the `settings/` directory to define the proxy server and the backend services.

```yaml
server:
  host: "localhost"
  listen_port: "8080"
  cert_file: ""   # Path to the SSL certificate (optional)
  key_file: ""    # Path to the SSL key (optional)

resources:
  - name: "server1"
    endpoint: "/server1/"
    destination_url: "http://localhost:9001"
  - name: "server2"
    endpoint: "/server2/"
    destination_url: "http://localhost:9002"
  - name: "server3"
    endpoint: "/server3/"
    destination_url: "http://localhost:9003"
```

- `host`: The host to bind the server to.
- `listen_port`: The port on which the proxy will listen.
- `cert_file`, `key_file`: Optional fields for SSL/TLS certificates.
- `resources`: List of backend services, with their endpoints and URLs.

### Running the Proxy Server

1. Start the backend services using Docker:

```bash
make run-containers
```

This will spin up three backend services (`server1`, `server2`, `server3`) using the `kennethreitz/httpbin` container.

2. Run the proxy server:

```bash
make run-proxy-server
```

The server will start listening on `http://localhost:8080`.

### Testing

- **Health Check**: Test the health check endpoint:

```bash
curl http://localhost:8080/ping
```

You should receive:

```plaintext
pong
```

- **Proxying Requests**: You can proxy requests to the backend services defined in your `config.yaml`. For example:

```bash
curl -I http://localhost:8080/server1
```

This will forward the request to `server1` running on `http://localhost:9001`.

### Stopping the Services

To stop the Docker containers:

```bash
make stop
```

## Caching

The proxy server caches responses for 30 seconds by default. The cache duration can be modified in the code where the cache is initialized:

```go
cache := NewCache(30 * time.Second) // Cache expiration time (30 seconds)
```

Cached responses are stored based on the request URL (including query parameters). Cached data is automatically evicted once it expires.

## Contributing

Feel free to open an issue or submit a pull request if you'd like to contribute to this project!
