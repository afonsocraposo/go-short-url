# go-short-url

## Overview

`go-short-url` is a URL shortener service written in Go (Golang). This project makes it easy to shorten URLs for easier sharing and management, while also offering retrieval of original URLs based on their shortened versions. It runs on a lightweight HTTP server on port 8080.

## Features

- GET and POST methods for shortening and retrieving URLs
- Customizable hash size for the generated URLs
- Basic error handling for unsupported methods and invalid URLs

## Prerequisites

- Go 1.16 or above

## Getting Started

To get started, clone the repository and navigate into its directory:

```bash
git clone https://github.com/yourusername/go-short-url.git
cd go-short-url
```

### Installing Dependencies

This project has no external dependencies.

### Building the Project

Compile the code:

```bash
go build .
```

### Running the Server

Execute the binary:

```bash
./go-short-url
```

If everything is set up correctly, you should see:

```
Server is running on port 8080
```

## API Endpoints

### Shorten a URL (`POST`)

To shorten a URL, make a POST request to the root `/`:

```bash
curl -d "url=http://example.com" -X POST http://127.0.0.1:8080/
```

### Retrieve a URL (`GET`)

To retrieve the original URL, navigate to the shortened URL in your browser or use a GET request. For example:

```bash
curl http://127.0.0.1:8080/{hash}
```

## Error Handling

- "Method not allowed": Unsupported HTTP method
- "Invalid URL format": URL doesn't meet the required format
- "URL not found": The hash doesn't correspond to any stored URL
- "Something went wrong": Internal server error

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is under the MIT License. See [LICENSE.md](LICENSE.md) for details.
