# Image Crawler & Harvester

A Go-based web crawler that discovers images on websites, processes them (metadata extraction and thumbnail generation), and stores the information in a MySQL database. It also provides a web interface for searching and viewing the harvested images.

## Features

- **Concurrent Crawling**: Efficiently crawls websites using a worker pool.
- **Image Processing**: Extracts image metadata (dimensions, format, alt text, title) and generates thumbnails.
- **Headless Browsing**: Uses `chromedp` to handle dynamic content (if applicable/needed by internal logic).
- **Persistence**: Stores image data in MySQL.
- **Web UI**: Searchable interface to filter images by dimensions, format, and metadata.

## Requirements

- **Go**: 1.25 or later.
- **MySQL**: 8.0 (can be run via Docker).
- **Chrome/Chromium**: Required by `chromedp` for crawling.

## Setup

### 1. Database

The project includes a `docker-compose.yml` to quickly spin up a MySQL instance.

```bash
docker-compose up -d
```

**Note**: You may need to manually create the `images` table if the application doesn't handle it automatically.
TODO: Add schema initialization script or instructions.

### 2. Dependencies

Install Go dependencies:

```bash
go mod download
```

## Running the Application

The application is executed via `main.go` and accepts various command-line flags.

### Basic Usage

```bash
go run . [options] <url1> <url2> ...
```

### Examples

```bash
# Crawl with default settings
go run . https://example.com

# Custom worker pool and timeout
go run . -workers=10 -max-goroutines=50 -timeout=5m https://example.com https://golang.org

# Follow external links
go run . -external https://example.com
```

### Command Line Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-workers` | `10` | Worker pool size. |
| `-max-goroutines` | `50` | Maximum number of concurrent goroutines. |
| `-external` | `false` | Whether to follow external links. |
| `-timeout` | `2m` | Total crawler timeout (e.g., `5m`, `1h`). |
| `-mysql` | `crawler:crawler@tcp(127.0.0.1:3306)/webcrawler` | MySQL Connection String (DSN). |

## Web Interface

Once the application starts, the Web UI is available at:

[http://localhost:8080](http://localhost:8080)

You can search for images using filters like:
- URL/Filename/Alt/Title (substring match)
- Format (e.g., `png`, `jpeg`)
- Minimum/Maximum Width and Height

## Project Structure

- `main.go`: Application entry point. Handles configuration, database connection, and starts the crawler and web server.
- `internal/config/`: Configuration loading and flag parsing.
- `internal/crawler/`: Core crawling logic, worker pool, and image parsing.
- `internal/storage/`: MySQL repository for image records.
- `internal/web/`: Web server and HTML templates.
- `assets/`: 
    - `originals/`: (TODO: Verify if used for original images)
    - `thumbnails/`: Generated image thumbnails.
- `docker-compose.yml`: Local database setup.

## Development

### Running Tests
TODO: No tests detected. Add unit tests for crawler and parser logic.

## License
Apache-2.0 license
