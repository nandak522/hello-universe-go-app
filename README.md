# hello-universe-go-app
hello-universe-go-app

# Run server
```sh
go build -o server && ./server
```

# Install as Helm chart
```sh
kubectl create ns hello-universe

cd charts/hello-universe

helm template -v 5 \
    --create-namespace \
    --namespace hello-universe \
    --logtostderr \
    --debug \
    --values values-default.yaml \
    hello-universe \
    .

helm upgrade -v 3 \
    --create-namespace \
    --namespace hello-universe \
    --logtostderr \
    --debug \
    --install \
    --atomic \
    --timeout 60s \
    --debug \
    --cleanup-on-fail \
    --values values-default.yaml \
    hello-universe \
    .
```

# Want to see Newrelic Instrumentation ?
Supply the below environment variables with valid values and start the service. Thats all.

* `NEW_RELIC_APP_NAME`
* `NEW_RELIC_LICENSE_KEY`

For example:
```sh
export NEW_RELIC_APP_NAME=hello-universe
export NEW_RELIC_LICENSE_KEY="YOUR VALID NEWRELIC INGEST LICENSE KEY"
go build -o server && ./server
```

Headover to Newrelic and find the application in "APM" screen.
