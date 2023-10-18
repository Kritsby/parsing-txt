FROM golang:1.21 AS builder

ENV CGO_ENABLED=0

ARG NAME=Profile

WORKDIR /src

COPY . .

RUN \
    go build -o /bin/app

FROM alpine:3.18

ARG APP_USER=app
ARG APP_UID=1010

RUN adduser -D -u $APP_UID $APP_USER

USER $APP_USER:$APP_USER
WORKDIR /srv

COPY --from=builder --chown=$APP_USER:$APP_USER /bin/app /srv/app

ENTRYPOINT ["/srv/app"]
