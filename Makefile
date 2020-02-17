build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/migrate handlers/migrate/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/workspace handlers/workspace/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/oauth handlers/oauth/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/mattermost handlers/mattermost/main.go

run:
	# build the package
	make build
	cp serverless.default.yaml serverless.yml
	# generate sam template for local development
	make run-local

run-local-server:
	# run sam server locally
	#sam local start-api
	sls offline --useDocker

run-local-db:
	# run database locally
	docker-compose -f  docker-compose.local.yaml up --build

run-local: run-local-db run-local-server
