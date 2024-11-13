# Pizza-UI

**Pizza-UI** is a reusable microservice for creating and serving UI components, leveraging [Templ](https://github.com/benbjohnson/templ) to generate templates. This project provides a server running on a specified port via the `UI_PORT` environment variable, offering a customizable user interface setup for other microservices, such as `auth`, `orders`, and more.

## Features

- Templ-based HTML component generation
- Environment-based configuration for project name and port
- Ready for containerization and cloud-native deployment
- Customizable and reusable for multiple applications

## Requirements

- **Go** (latest stable version recommended)
- **Docker** (optional for containerized deployment)

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/username/pizza-ui.git
cd pizza-ui

### 2. Initial Configuration

To set up the project, define the project name and the port on which it will run. You can configure this through a `.env` file.

1. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```

2. Edit `.env` to set the `PROJECT_NAME` and `UI_PORT` environment variables:

   ```plaintext
   PROJECT_NAME=pizza-ui
   UI_PORT=8080
   ```

   Alternatively, you can directly set these variables in your environment.

### 3. Running the Server

To start the server, run:

```bash
go run main.go
```

This will start the server on the port defined by `UI_PORT`. If you haven't set `UI_PORT`, the server defaults to port `8080`.

### 4. Accessing the Application

Once the server is running, open your web browser and go to:

```
http://localhost:8080
```

Or use the port you specified in `UI_PORT`.

### 5. Dockerized Deployment (Optional)

To run the Pizza-UI service as a Docker container:

1. **Build the Docker Image:**
   ```bash
   docker build -t pizza-ui .
   ```

2. **Run the Docker Container:**
   ```bash
   docker run -p 8080:8080 -e UI_PORT=8080 pizza-ui
   ```

   Adjust the `UI_PORT` variable as needed.

### 6. Testing

For testing, use:

```bash
go test ./...
```
