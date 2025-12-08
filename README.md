# Echo Server

A simple Go web server that echoes back query parameters.

## Running the Server

### Option 1: Run with Go

```bash
go run main.go
```

### Option 2: Run with Docker

Build the Docker image:
```bash
docker build -t echo-server .
```

Run the container:
```bash
docker run -p 8080:8080 echo-server
```

### Option 3: Run with Docker Compose

Start the service:
```bash
docker-compose up
```

Stop the service:
```bash
docker-compose down
```

The server will start on `http://localhost:8080`

## Usage

Send HTTP requests with query parameters:

```bash
curl "http://localhost:8080/?name=John&age=30"
```

Output:
```
Echo - Query Parameters:
========================

name: John
age: 30
```

## Examples

```bash
# Single parameter
curl "http://localhost:8080/?message=Hello"

# Multiple parameters
curl "http://localhost:8080/?foo=bar&baz=qux"

# Multiple values for same parameter
curl "http://localhost:8080/?color=red&color=blue"
```
