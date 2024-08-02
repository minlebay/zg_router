--- 

# Router Component

The Router component is a part of the ZmeyGorynych Project. It receives messages from the generator and routes them to various processing servers using gRPC.

## Components

### Router (`zg_router`)
This component routes messages received from the generator to the processing servers.

#### Docker Compose Configuration
```yaml
version: '3.8'

networks:
  local-net:
    external: true

services:
  zg_router:
    build:
      context: .
      dockerfile: ./Dockerfile
    container_name: zg_router
    env_file:
      - .env-docker
    networks:
      - local-net
    ports:
      - "21123:21123"
    volumes:
      - .:/app
    restart: unless-stopped
```

#### Configuration File (`.config.yaml`)
```yaml
prometheus:
  url: ${PROMETHEUS_URL}

grpc_server:
  listen_address: ${GRPC_SERVER_LISTEN_ADDRESS}

processing:
  processing_servers_list:
    - ${GRPC_SERVER_PROCESSING_ADDRESS}

logstash:
  url: ${LOGSTASH_URL}
```

#### .env-docker File
```env
GRPC_SERVER_LISTEN_ADDRESS=zg_router:50051
GRPC_SERVER_PROCESSING_ADDRESS=zg_processing:50052
LOGSTASH_URL=http://logstash:5000
PROMETHEUS_URL=0.0.0.0:21123
```

## Other Components

- **Message Generator**: Generates messages and sends them to the router.
- **Processing Servers**: Multiple servers that process the received messages.
- **Prometheus**: Monitors the application and collects metrics.
- **ELK Stack**: Collects and analyzes logs.
- **Grafana**: Visualizes the metrics collected by Prometheus.
- **Kafka**: A message broker that integrates with the backend.
- **Databases**: Includes MongoDB, MySQL, Redis for caching and indexing, and SQL/NoSQL repositories.

## Getting Started

### Prerequisites
- Docker
- Docker Compose

### Running the Router
1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/message-generator.git
   cd message-generator/router
   ```
2. Build and run the Docker containers:
   ```bash
   docker-compose up --build
   ```

### Environment Variables
Ensure to set the following environment variables in the `.env-docker` file:
- `GRPC_SERVER_LISTEN_ADDRESS`: Address of the gRPC server (e.g., `zg_router:50051`).
- `GRPC_SERVER_PROCESSING_ADDRESS`: Address of the processing server (e.g., `zg_processing:50052`).
- `LOGSTASH_URL`: URL of the Logstash server (e.g., `http://logstash:5000`).
- `PROMETHEUS_URL`: URL of the Prometheus server (e.g., `0.0.0.0:21123`).

## Contributing
Contributions are welcome! Please fork the repository and create a pull request with your changes.

## License
This project is licensed under the MIT License.

---