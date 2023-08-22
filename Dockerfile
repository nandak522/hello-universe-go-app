FROM golang:1.21 as builder
ENV USER=appuser \
    UID=10001 \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"
WORKDIR /app
RUN mkdir -p /app/templates
RUN mkdir -p /app/static
COPY *.go /app/
COPY go.* /app/
COPY templates /app/templates/
COPY static /app/static/
RUN cd /app && \
    go build -ldflags="-s -w" -a -o server

FROM ubuntu:23.04
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
ENV DEBIAN_FRONTEND=noninteractive \
    USER=appuser \
    APP_PORT=8000
WORKDIR /app
COPY --chown=${USER}:${USER} --from=builder /app/server .
RUN chown -Rv ${USER}:${USER} /app
EXPOSE ${APP_PORT}
RUN apt-get -qq update && apt-get -qq install -y --no-install-recommends curl net-tools netcat-openbsd dnsutils && \
    apt clean && \
    rm -rf /var/lib/apt/lists/*
USER ${USER}:${USER}
CMD ["/app/server"]
