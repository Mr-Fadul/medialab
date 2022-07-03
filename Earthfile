VERSION 0.6

all-unit-test:
    BUILD ./services/digitaltv-recorder+unit-test

all-docker:
    BUILD ./services/digitaltv-recorder+docker

all-release:
    BUILD ./services/digitaltv-recorder+release

dev-up:
    LOCALLY
    RUN docker-compose -f deployments/docker-compose/docker-compose.yaml up

dev-down:
    LOCALLY
    RUN docker-compose -f deployments/docker-compose/docker-compose.yaml down