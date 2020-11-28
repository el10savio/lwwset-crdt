
provision:
	@echo "Provisioning LWWSet Cluster"	
	bash scripts/provision.sh

lwwset-build:
	@echo "Building LWWSet Docker Image"	
	docker build -t lwwset -f Dockerfile .

lwwset-run:
	@echo "Running Single LWWSet Docker Container"
	docker run -p 8080:8080 -d lwwset

info:
	echo "LWWSet Cluster Nodes"
	docker ps | grep 'lwwset'
	docker network ls | grep lwwset_network

clean:
	@echo "Cleaning LWWSet Cluster"
	docker ps -a | awk '$$2 ~ /lwwset/ {print $$1}' | xargs -I {} docker rm -f {}
	docker network rm lwwset_network

build:
	@echo "Building LWWSet Server"	
	go build -o bin/lwwset main.go

fmt:
	@echo "go fmt LWWSet Server"	
	go fmt ./...

test:
	@echo "Testing LWWSet"	
	go test -v --cover ./...