build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/migrate handlers/migrate/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/workspace handlers/workspace/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/oauth handlers/oauth/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/mattermost handlers/mattermost/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/project handlers/project/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/bot handlers/bot/main.go

run:
	# build the package
	make build
	cp serverless.default.yaml serverless.yml
	# generate sam template for local development
	make run-local

run-local-server:
	# run sam server locally
	#sam local start-api
	sls offline --useDocker --host $(DATABASE_HOSTNAME)

migrate-local:
	sls invoke local -f migrate --stage dev

run-local-db:
	# run database locally
	docker-compose -f  docker-compose.local.yaml up --build

run-local: run-local-db run-local-server migrate-local

deploy:
	make build

	cp serverless.default.yaml serverless.yml
	sls deploy --stage $(STAGE) -v --region $(REGION)

	make migrate

remove:
	sls remove --stage $(STAGE) -v --region $(REGION)

migrate:
	sls invoke -f migrate --stage $(STAGE) -v --region $(REGION)