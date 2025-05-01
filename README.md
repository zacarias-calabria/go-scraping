# Go Scraper

A web scraping application written in Go that allows efficient and configurable extraction of information from websites.

## Features

- Asynchronous web scraping
- Configurable rate limiting
- Automatic User-Agent rotation
- Proxy support
- Timeout control
- Configurable link depth tracking
- Error handling and HTTP status codes management

## Requirements

- Go 1.24 or higher
- Docker (optional, for containerized execution)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/your-username/go-scraping.git
cd go-scraping
```

2. Install dependencies:
```bash
go mod download
```

## Usage

The application can be run directly or using Docker.

### Direct Execution

```bash
go run main.go -url https://example.com [-timeout 30s] [-proxy http://proxy:port]
```

Parameters:
- `-url`: URL of the site to scrape (required)
- `-timeout`: Maximum wait time (default: 30s)
- `-proxy`: Proxy URL to use (optional)

### Docker Execution

1. Build the image:
```bash
docker compose build
```

2. Start the container:
```bash
docker compose up -d
```

3. Run the scraper inside the container:
```bash
docker compose exec app go run main.go -url https://example.com
```

To stop the container:
```bash
docker compose down
```

## Project Structure

```
.
├── internal/
│   ├── adapters/     # Adapter implementations
│   ├── domain/       # Domain models
│   └── ports/        # Port interfaces
├── main.go          # Application entry point
├── go.mod           # Dependencies file
├── go.sum           # Dependencies checksum
├── Dockerfile       # Docker configuration
└── compose.yaml     # Docker Compose configuration
```

## Development

### Testing

To run tests:
```bash
go test ./...
```

### Debugging

The application is configured for debugging with Delve. To debug:

1. Start the debug server:
```bash
dlv debug
```

2. Connect to port 40000 from your IDE.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
