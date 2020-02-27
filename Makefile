build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/migrate handlers/migrate/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/workspace handlers/workspace/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/oauth handlers/oauth/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/mattermost handlers/mattermost/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/project handlers/project/main.go

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

run-local-db:
	# run database locally
	docker-compose -f  docker-compose.local.yaml up --build
#   echo '${DATABASE_HOSTNAME}'

run-local: run-local-db run-local-server

deploy:
	make build

	cp serverless.default.yaml serverless.yml

#	ifdef prod
#    sls deploy --stage prod -v --region us-east-1
#    else
#	ifdef stage
#	sls deploy --stage stage -v --region ap-south-1
#    else
#    sls deploy --stage dev -v --region ap-south-1
#    endif
	sls deploy --stage stage -v --region ap-south-1



remove:
	ifdef prod
	sls remove --stage prod -v --region us-east-1
	else
	ifdef stage
	sls remove --stage stage -v --region ap-south-1
	else
	sls remove --stage dev -v --region ap-south-1
	endif