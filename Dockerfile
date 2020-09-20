FROM golang:1.15 as go-builder
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /app
COPY go.mod /app
COPY go.sum /app
COPY cmd /app/cmd/
COPY pkg /app/pkg/
RUN go build -o server /app/cmd/server && chmod +x /app/server

FROM scratch
COPY --from=go-builder /app/server .
EXPOSE 8888
ENTRYPOINT ["/server"]
