VERSION 0.6

all-unit-test:
    BUILD ./pkg/utils+unit-test
    BUILD ./services/digitaltv-recorder+unit-test
    BUILD ./services/handle-video-intelligence+unit-test

all-docker:
    BUILD ./services/digitaltv-recorder+docker
    BUILD ./services/handle-video-intelligence+docker

all-release:
    BUILD ./services/digitaltv-recorder+release
    BUILD ./services/handle-video-intelligence+release

# docker-compose section
digital-recorder-dev-up:
    LOCALLY
    RUN docker-compose -f deployments/docker-compose/digital-recorder/docker-compose.yaml up

digital-recorder-dev-down:
    LOCALLY
    RUN docker-compose -f deployments/docker-compose/digital-recorder/docker-compose.yaml down

cloud-run-dev-up:
    LOCALLY
    RUN docker-compose -f deployments/docker-compose/cloud-run/docker-compose.yaml up

cloud-run-dev-down:
    LOCALLY
    RUN docker-compose -f deployments/docker-compose/cloud-run/docker-compose.yaml down