build:
	docker-compose -f .\deployments\docker-compose.yaml build
up: 
	docker-compose -f .\deployments\docker-compose.yaml up
down:
	docker-compose -f .\deployments\docker-compose.yaml down

store:
	docker create --name postgres-test -p 5432:5432 -e POSTGRES_PASSWORD=123 -e POSTGRES_DB=jetstyle_test postgres
test:
	docker start postgres-test
	go test -v -timeout 30s ./tests/...
	docker stop postgres-test

