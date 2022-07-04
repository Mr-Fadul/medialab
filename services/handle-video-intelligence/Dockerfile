FROM ubuntu:22.04

ARG app_version

ENV APP_VERSION $app_version
ENV DEBIAN_FRONTEND=noninteractive

ARG app_version

ENV APP_VERSION $app_version

RUN apt-get update && \
    apt-get install -y \
    tzdata \
    dumb-init \
    curl \
    iputils-ping \
    dumb-init \ 
    ca-certificates \
    && update-ca-certificates 2>/dev/null || true

RUN rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

WORKDIR /app
COPY handle-video-intelligence /app/bin/

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/app/bin/handle-video-intelligence"]