APP_VERSION=edge
APP_NAME=callofduty
BIN_DIR=bin

init:
	@mkdir -p $(BIN_DIR)
	@go mod download

run-devserver: init
	@buffalo dev

build: init
	@buffalo build --static -o $(BIN_DIR)/$(APP_NAME)

run: init
	$(BIN_DIR)/$(APP_NAME)

build-docker-image:
	@echo "Building $(APP_NAME):$(APP_VERSION) docker image ..."
	@docker build -t ashokrajar/$(APP_NAME):$(APP_VERSION) -t $(APP_NAME):$(APP_VERSION) .

push-docker-image:
	@echo "Pushing $(APP_NAME):$(APP_VERSION) image to Docker Hub ..."
	@docker push ashokrajar/$(APP_NAME):$(APP_VERSION)

build-push-docker-image: build-docker-image push-docker-image

run-in-docker: build-docker-image
	@echo "Running $(APP_NAME):$(APP_VERSION) locally using docker-compose ..."
	@docker-compose up
