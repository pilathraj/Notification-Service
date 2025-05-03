# Setup 

```cmd
1. curl -sSL \
https://raw.githubusercontent.com/bitnami/containers/main/bitnami/kafka/docker-compose.yml > docker-compose.yml

2. Find KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://:9092
And replace it with:
KAFKA_CFG_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092
3. docker pull docker.io/bitnami/kafka:4.0  &&  docker-compose up -d
4. go mod init Notification-Service
5. go get github.com/IBM/sarama github.com/gin-gonic/gin
6. go run cmd/main/main.go
7. go run 
8. curl -X POST http://localhost:8081/send \
-d "fromID=2&toID=1&message=Prakash started following you."
curl -X POST http://localhost:8081/send \
-d "fromID=4&toID=1&message=Manasa liked your post: 'My weekend getaway!'"
curl http://localhost:8081/notifications/1
```

```cmd
minikube image load notification-app:latest
```