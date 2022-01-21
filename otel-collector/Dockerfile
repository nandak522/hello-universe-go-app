FROM golang:1.17 as builder
WORKDIR /app
COPY otelcol /app/
COPY otel-config.yaml /app/
RUN chmod +x /app/otelcol

FROM scratch
WORKDIR /app
COPY --from=builder /tmp /tmp
COPY --from=builder /app/otelcol /app/otelcol
COPY --from=builder /app/otel-config.yaml /app/otel-config.yaml
EXPOSE 4317
CMD ["/app/otelcol", "--config", "/app/otel-config.yaml" ]
