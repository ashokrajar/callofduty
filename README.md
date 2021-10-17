# Call Of Duty

**Call Of Duty** is a simple SRE On-Call handover tool.

## Quick Starter

### Prerequisite
- Docker

```shell
mkdir data

export DB_PASSWORD="*********" # replace with your password of choice

docker run --name callofduty-prod-db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=${DB_PASSWORD} -e POSTGRES_DB=callofduty_production -p 5432:5432 -v data:/var/lib/postgresql/data -d postgres:14.0-alpine3.14

docker run --name callofduty-prod -e GO_ENV=production -e DATABASE_HOST=localhost -e DATABASE_PASSWORD=${DB_PASSWORD} -p 8080:3000 -d ashokrajar/callofduty:0.0.1 

```
Go to => [http://localhost:8080](http://localhost:8080)

### Running with Docker Compose

```shell
export GO_ENV="production"
export APP_VERSION="0.0.1"
export APP_DBPASS="******" # set your choice of DB password

git clone https://github.com/ashokrajar/callofduty.git && cd callofduty
make run-in-docker
```
Go to => [http://localhost:8080](http://localhost:8080)

### Building and Running from Source

You will need a locally running PostgreSQL server.

```shell
export GO_ENV="development"
export DATABASE_HOST="locahost" # your database user
export DATABASE_USER="postgres" # your database user
export DATABASE_PASSWORD="******" # your database password
git clone https://github.com/ashokrajar/callofduty.git && cd callofduty 
make run
```
Go to => [http://localhost:3000](http://localhost:3000)

## Local Development Getting Started

### Prerequisite
- make
- GoLang 1.17.x
- Docker
- Docker Compose

### Running DevServer

Running a DevServer is useful as it applies and reloads the development webserver as you write code.

Dependency : You will need a locally running PostgreSQL server.

```shell
export GO_ENV="development"
export DATABASE_HOST="locahost" # your database user
export DATABASE_USER="postgres" # your database user
export DATABASE_PASSWORD="******" # your database password

make run-devserver
```

## Deployment with K8

**Work In Progress**

[Powered by Buffalo](http://gobuffalo.io)
