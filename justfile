alias cu:= compose-up
alias cd:= compose-down

# Start the Docker Compose services in detached mode
compose-up:
	docker-compose -f infra/dev/docker-compose.yaml up -d

# Stop and remove containers, networks, images, and volumes
compose-down:
	docker-compose -f infra/dev/docker-compose.yaml down
