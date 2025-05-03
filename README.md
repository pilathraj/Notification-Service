# Notification-Service

Notification-Service is a microservice designed to handle notifications. It uses Kafka for messaging and Docker for containerization.

## Prerequisites

- Docker
- Docker Compose
- Minikube (optional, for Kubernetes deployment)
- Go (for building the project)

## Makefile Commands

The `Makefile` provides the following commands:

- `make init`: Initializes the project.
- `make install`: Installs dependencies using `go get`.
- `make build`: Builds the Docker image for the project.
- `make deploy`: Deploys the project using `docker-compose`.
- `make deploy-minikube`: Deploys the project to Minikube (requires Kubernetes setup).
- `make stop`: Stops the running Docker containers.
- `make all`: Runs `init`, `install`, `build`, and `deploy` in sequence.

## Docker Compose Services

The `docker-compose.yml` defines the following services:

1. **Zookeeper**:
    - Image: `bitnami/zookeeper:latest`
    - Ports: `2181:2181`
    - Environment: `ALLOW_ANONYMOUS_LOGIN=yes`

2. **Kafka**:
    - Image: `bitnami/kafka:latest`
    - Ports: `9092:9092`
    - Environment:
      - `KAFKA_BROKER_ID=1`
      - `KAFKA_CFG_ZOOKEEPER_CONNECT=zookeeper:2181`
      - `KAFKA_CFG_LISTENERS=PLAINTEXT://:9092`
      - `KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://kafka:9092`
      - `ALLOW_PLAINTEXT_LISTENER=yes`
    - Depends on: `zookeeper`

3. **Notification-App**:
    - Image: `notification-app:latest`
    - Ports: `8081:8081`
    - Environment: `KAFKA_BROKERS=kafka:9092`
    - Depends on: `kafka`

## Usage

### Build and Deploy

To build and deploy the project, follow these steps:

1. Initialize the project:
    ```bash
    make init
    ```

2. Install project dependencies:
    ```bash
    make install
    ```

3. Build the Docker image:
    ```bash
    make build
    ```

4. Deploy the project using `docker-compose`:
    ```bash
    make deploy
    ```

5. (Optional) Deploy the project to Minikube:
    ```bash
    make deploy-minikube
    ```

6. Stop the running Docker containers:
    ```bash
    make stop
    ```

7. Run all the above steps in sequence:
    ```bash
    make all
    ```
