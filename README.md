## Pre-requisite 

- [Node](https://nodejs.org/en/download/) 
- [Golang](https://golang.org/dl/)

- Serverless Framework 

```bash
npm install -g serverless
```

- [Docker](https://www.docker.com/products/docker-desktop)

## Run Application

```bash

yarn install

export DATABASE_NAME='goscrum'
export DATABASE_HOSTNAME='192.168.31.57' ## IP address of your machine
export DATABASE_USERNAME='goscrum'
export DATABASE_PASSWORD='goscrum'

make run -j
```