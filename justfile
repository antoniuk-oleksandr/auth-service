alias cu:= compose-up
alias cd:= compose-down
alias cub:= compose-up-build
alias mongo:= compose-up-mongo

# Start the Docker Compose services in detached mode
compose-up:
	docker-compose -f infra/dev/docker-compose.yaml up -d

# Stop and remove containers, networks, images, and volumes
compose-down:
	docker-compose -f infra/dev/docker-compose.yaml down

# Rebuild the Docker Compose services and start them in detached mode
compose-up-build:
    docker-compose -f infra/dev/docker-compose.yaml up -d --build

# Start only the MongoDB services in detached mode
compose-up-mongo:
    docker-compose -f infra/dev/docker-compose.yaml up -d ctp-mongo ctp-mongo-setup