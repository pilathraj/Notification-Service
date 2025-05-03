.PHONY: init install build deploy deploy-minikube stop all

init:
	@echo "Initializing..."

install:
	@echo "Installing dependencies..."
	go get -u ./...


build:
	@echo "Building the project..."
	docker build -t notification-service:latest .

deploy:
	@echo "Deploying the project..."
	docker-compose up -d

minikube-start:
	@echo "Starting Minikube..."
	minikube start --driver=docker
	

deploy-minikube:
	@echo "Deploying to Minikube..."
	kubectl apply -f kafka-deployment.yaml
	kubectl apply -f kafka-service.yaml
	kubectl apply -f notification-app-deployment.yaml
	kubectl apply -f notification-app-service.yaml
	
minikube-stop:
	@echo "Stopping Minikube..."
	minikube stop
	minikube delete

minikube-status:
	@echo "Checking Minikube status..."
	kubectl get deployments
	kubectl get services

stop:
	@echo "Stopping the project..."
	docker-compose down

all: init install build deploy
	@echo "All tasks completed successfully."

debug:
	@echo "Debugging the project..."
	#docker-compose up --build
	kubectl get pods
	# kubectl describe pod <pod-name>
	# kubectl logs -f <pod-name>
	kubectl get service notification-app-service
	# kubectl logs <YOUR_NOTIFICATION_APP_POD_NAME>
	# minikube image ls 
	# kubectl delete -f notification-app-deployment.yaml and apply again

