FROM bash:latest

RUN apk add --no-cache --upgrade grep && \
    apk add --no-cache gettext curl wget jq
