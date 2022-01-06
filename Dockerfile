FROM golang:1.17 as builder
ENV USER=appuser \
    UID=10001 \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
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
RUN cd /app \
    && go build -a -o server

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
WORKDIR /app
USER ${USER}:${USER}
COPY --from=builder /app/server .
ENV USER=appuser \
    APP_PORT=8000
EXPOSE ${APP_PORT}
CMD ["/app/server"]
