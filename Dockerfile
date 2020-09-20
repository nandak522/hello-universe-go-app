FROM golang:1.15 as builder
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
COPY *.go /app/
COPY go.* /app/
COPY templates /app/templates/
COPY go.* /app/
RUN cd /app \
    && go build -a -o server

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
WORKDIR /app
USER ${USER}:${USER}
COPY --from=builder /app/server .
COPY --from=builder /app/templates /app/templates/
ENV USER=appuser \
    APP_PORT=1323
EXPOSE ${APP_PORT}
CMD ["/app/server"]
