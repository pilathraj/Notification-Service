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
	

minikube-deploy:
	@echo "Deploying to Minikube..."
	kubectl apply -f kafka-deployment.yaml
	kubectl apply -f kafka-service.yaml
	kubectl apply -f notification-app-deployment.yaml
	kubectl apply -f notification-app-service.yaml
	
minikube-stop:
	@echo "Stopping Minikube..."
	minikube stop
	minikube delete

minikube-stop-services:
	@echo "Stopping Minikube..."
	kubectl delete -f kafka-deployment.yaml
	kubectl delete -f kafka-service.yaml
	kubectl delete -f notification-app-deployment.yaml
	kubectl delete -f notification-app-service.yaml

minikube-status:
	@echo "Checking Minikube status..."
	kubectl get deployments
	kubectl get services

minikube-tunnel:
	@echo "Starting Minikube tunnel..."
	minikube tunnel

stop:
	@echo "Stopping the project..."
	docker-compose down

all: init install build deploy
	@echo "All tasks completed successfully."

minikube-debug:
	@echo "Debugging the project..."
	#kubectl get pods
	kubectl get pods
	@echo "---------------------------------------------------------"
	# kubectl describe pod <pod-name>
	kubectl describe pod $(shell kubectl get pods -l app=notification-app -o jsonpath="{.items[0].metadata.name}")
	@echo "---------------------------------------------------------"
	# kubectl get service
	kubectl get service kafka-service
	kubectl get service notification-app-service
	@echo "---------------------------------------------------------"
	# kubectl logs -f <YOUR_NOTIFICATION_APP_POD_NAME>
	kubectl logs -f $(shell kubectl get pods -l app=notification-app -o jsonpath="{.items[0].metadata.name}")
	@echo "---------------------------------------------------------"
	# minikube image ls 

