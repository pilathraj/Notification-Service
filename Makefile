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
	minikube start --driver=docker --cpus=4 --memory=4200 --disk-size=20g

minikube-deploy-infrastructure:
	@echo "Deploying to Minikube..."
	# Apply PersistentVolumes and PersistentVolumeClaims
	kubectl apply -f k8s/postgres-pv.yaml
	kubectl apply -f k8s/kafka-pv.yaml
	# Apply ConfigMaps
	kubectl apply -f k8s/postgres-configmap.yaml
	kubectl apply -f k8s/kafka-configmap.yaml
	# Apply Deployments
	kubectl apply -f k8s/postgres-deployment.yaml
	kubectl apply -f k8s/kafka-deployment.yaml
	kubectl apply -f k8s/notification-app-deployment.yaml
	# Apply Services
	kubectl apply -f k8s/postgres-service.yaml
	kubectl apply -f k8s/kafka-service.yaml
	kubectl apply -f k8s/notification-app-service.yaml


minikube-deploy:
	@echo "Deploying notification app to Minikube..."
	kubectl apply -f k8s/notification-app-deployment.yaml
	kubectl apply -f k8s/notification-app-service.yaml
	# Apply cronjob
	kubectl apply -f k8s/kafka-consumer-cronjob.yaml

	
minikube-stop:
	@echo "Stopping Minikube..."
	minikube stop
	minikube delete

minikube-stop-services:
	@echo "Stopping Minikube..."
	kubectl delete -f k8s/postgres-deployment.yaml
	kubectl delete -f k8s/postgres-service.yaml
	kubectl delete -f k8s/kafka-deployment.yaml
	kubectl delete -f k8s/kafka-service.yaml
	kubectl delete -f k8s/notification-app-deployment.yaml
	kubectl delete -f k8s/notification-app-service.yaml
	kubectl delete -f k8s/kafka-consumer-cronjob.yaml

minikube-status:
	@echo "Checking Minikube status..."
	kubectl get pv
	kubectl get pvc
	kubectl get deployments
	kubectl get services
	kubectl get pods

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

