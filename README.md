# Docker Lab Echo

This Docker lab demonstrates building a Docker image, and running it with Docker build
and Docker compose.  A simple Go web server that echoes back query parameters is used
for the demonstration.  

## Running the Echo Server

The echo server can be run using `docker run` or `docker compose`.  For this lab, the 
advantage to using `docker compose` is that it encodes all the options to run the server.

### Option 1: Run with Docker

Build the Docker image, and tag it as 'echo-server':
```bash
docker build -t echo-server .
```

Run the container, publish the container's port 8080 to the host's port 8080, and 
detach from the container:
```bash
docker run -d -p 8080:8080 echo-server
```

To stop the container use `docker ps` to find the ID or name for the running container,
and use `docker stop` to stop the container.  For example:
```bash
% docker ps
CONTAINER ID   IMAGE         COMMAND           CREATED         STATUS         PORTS                                         NAMES
428e9ccd9802   echo-server   "./echo-server"   3 seconds ago   Up 2 seconds   0.0.0.0:8080->8080/tcp, [::]:8080->8080/tcp   modest_visvesvaraya
% docker stop modest_visvesvaraya
```

### Option 2: Run with Docker Compose

Build the image:
```bash
docker compose build
```

Start the server:
```bash
docker compose up
```

Stop the server:
```bash
docker compose down
```

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
