Notification Service
# Notification Service

This project is a notification service that uses Kafka for message brokering. It includes a Kafka service and a notification-app service, which are managed using Docker Compose or deployed to Minikube for Kubernetes-based environments.

## Prerequisites
- Docker
- Docker Compose
- Minikube
- Go (for development)
- Kafka:
Ensure Kafka is running locally on localhost:9092.
Create the notifications topic:
```cmd 
kafka-topics --create --topic notifications --bootstrap-server localhost:9092
or
docker exec -it kafka-service /opt/bitnami/kafka/bin/kafka-topics.sh --create --topic notifications --bootstrap-server localhost:9092
```
- create topic inside minikube service
```cmd
/opt/bitnami/kafka/bin/kafka-topics.sh --create --topic notifications --bootstrap-server localhost:9092

- PostgreSQL:
Ensure PostgreSQL is running locally on localhost:5432.
Create a database named notifications:
```cmd
CREATE DATABASE notifications_db;
```


## Project Structure
### Services
#### Kafka Service:
- Uses the `bitnami/kafka` image.
- Configured with multiple listeners:
    - `PLAINTEXT://kafka-service:9092` for internal communication.
    - `EXTERNAL://localhost:9094` for external communication.
- Health checks are included to ensure Kafka is running properly.

#### Notification App:
- A custom application that connects to Kafka as a producer/consumer.
- Exposes port `8081` for external communication.
- Configured with the following environment variables:
    - `KAFKA_BROKERS`: Kafka broker address (`kafka:9092`).
    - `KAFKA_BOOTSTRAP_SERVERS`: Kafka bootstrap servers (`kafka:9092`).
    - `KAFKA_CONSUMER_GROUP`: Consumer group for the app (`notification-app`).

## Usage
### Local Development with Docker Compose
1. Initialize the Project:
     ```
     make init
     ```

2. Install Dependencies:
     ```
     make install
     ```

3. Build the Docker Image:
     ```
     make build
     ```

4. Deploy the Services:
     ```
     make deploy
     ```

5. Stop the Services:
     ```
     make stop
     ```

### Deployment to Minikube
1. Start Minikube:
     ```
     make minikube-start
     ```

2. Deploy to Minikube:
     ```
     make minikube-deploy
     ```

3. Check Minikube Status:
     ```
     make minikube-status
     ```

4. Start Minikube Tunnel (for LoadBalancer services):
     ```
     make minikube-tunnel
     ```

5. Stop Minikube Services:
     ```
     make minikube-stop-services
     ```

6. Stop and Delete Minikube:
     ```
     make minikube-stop
     ```

### Debugging in Minikube
To debug the application in Minikube, use:
```
make minikube-debug
```

This will:
- List all pods.
- Describe the notification-app pod.
- Display the services (kafka-service and notification-app-service).
- Tail the logs of the notification-app pod.

## Configuration
### Kafka Service
The Kafka service is configured with the following:
- Listeners:
    - `PLAINTEXT://kafka-service:9092` (internal communication).
    - `EXTERNAL://localhost:9094` (external communication).
- Environment Variables:
    - `KAFKA_CFG_NODE_ID=0`
    - `KAFKA_CFG_PROCESS_ROLES=controller,broker`
    - `KAFKA_CFG_CONTROLLER_QUORUM_VOTERS=0@kafka-service:9093`
    - `KAFKA_CFG_CONTROLLER_LISTENER_NAMES=CONTROLLER`
    - `KAFKA_CFG_LISTENER_SECURITY_PROTOCOL_MAP=CONTROLLER:PLAINTEXT,EXTERNAL:PLAINTEXT,PLAINTEXT:PLAINTEXT`

### Notification App
The notification-app is configured with the following environment variables:
- `KAFKA_BROKERS=kafka:9092`
- `KAFKA_BOOTSTRAP_SERVERS=kafka:9092`
- `KAFKA_CONSUMER_GROUP=notification-app`

## Volumes
### Kafka Data
A local volume (`kafka_data`) is used to persist Kafka data.

## Makefile Commands
| Command | Description |
|---------|-------------|
| `make init` | Initializes the project. |
| `make install` | Installs dependencies. |
| `make build` | Builds the Docker image for the notification app. |
| `make deploy` | Deploys the services using Docker Compose. |
| `make stop` | Stops the services. |
| `make minikube-start` | Starts Minikube. |
| `make minikube-deploy` | Deploys the services to Minikube. |
| `make minikube-status` | Checks the status of Minikube deployments. |
| `make minikube-tunnel` | Starts a Minikube tunnel for LoadBalancer services. |
| `make minikube-stop` | Stops and deletes Minikube. |
| `make minikube-stop-services` | Stops the Minikube services. |
| `make minikube-debug` | Debugs the Minikube deployment. |

## Notes
- Ensure that Docker and Minikube are installed and running before executing the commands.
- Use `make all` to initialize, install dependencies, build, and deploy the project in one step.
