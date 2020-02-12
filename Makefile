build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/migrate handlers/migrate/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/workspace handlers/workspace/main.go
	sls package

run:
	# build the package
	make build
	cp serverless.default.yaml serverless.yml
	# generate sam template for local development
	sls package
	sls sam export --output ./template.yml
	make run-local

run-local-server:
	# run sam server locally
	sam local start-api

run-local-db:
	# run database locally
	docker-compose -f  docker-compose.local.yaml up --build

run-local: run-local-db run-local-server
