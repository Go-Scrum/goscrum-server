## Pre-requisite 

- [Node](https://nodejs.org/en/download/) 
- [Golang](https://golang.org/dl/)

- Serverless Framework 

```bash
npm install -g serverless
```

- [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install-mac.html)

```bash
brew tap aws/tap
brew install aws-sam-cli
```

- [Docker](https://www.docker.com/products/docker-desktop)

## Run Application

```bash

export DATABASE_NAME='goscrum'
export DATABASE_HOSTNAME='192.168.31.56' ## IP address of your machine
export DATABASE_USERNAME='goscrum'
export DATABASE_PASSWORD='goscrum'

make run -j
```