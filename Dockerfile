FROM golang:1.14.6 as builder
RUN mkdir -p /app/templates
WORKDIR /app
COPY *.go /app/
COPY templates /app/templates/
COPY go.* /app/
RUN cd /app \
    && go mod tidy \
    && go build -o server .

FROM ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive \
    TZ=Asia/Kolkata \
    DEBUG_DEPS="curl less lsof strace netcat net-tools" \
    BUILD_DEPS="build-essential" \
    APP_DEPS="tzdata" \
    APP_USER=appuser \
    APP_PORT=1323

WORKDIR /app
RUN mkdir /app/templates
COPY --from=builder /app/server /app/
COPY --from=builder /app/templates /app/templates/

RUN set -ex \
    && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && groupadd --system ${APP_USER} \
    && useradd --no-log-init --system --create-home --gid ${APP_USER} ${APP_USER} \
    && usermod -u 1000 ${APP_USER} \
    && groupmod -g 1000 ${APP_USER} \
    && chown -Rv ${APP_USER}:${APP_USER} /app \
    && apt-get update && apt-get install -y --no-install-recommends ${BUILD_DEPS} ${APP_DEPS} ${DEBUG_DEPS} \
    && rm -rf /usr/share/doc && rm -rf /usr/share/man \
    && apt-get purge -y --auto-remove -o APT::AutoRemove::RecommendsImportant=false ${BUILD_DEPS} \
    && apt-get autoremove -y \
    && rm -rf /var/lib/apt/lists/*

USER ${APP_USER}
EXPOSE 1323
CMD ["/app/server"]
