# hello-universe-go-app
`hello-universe-go-app`, a simple "Hello-Universe" service written in Go, with built-in Newrelic instrumentation.

---
## Run server
```sh
go build -o server && ./server
```

---
## Install as Helm chart
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

---
## Want to see Newrelic Instrumentation ?
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

---
## Create a Git Tag
Whenever a new release/tag has to be created, just update `version.go` and push it to `main` branch. A github workflow is already configured that creates the actual (git) tag, which will be available in https://github.com/none-da/hello-universe-go-app/tags page.

> For now Github release creation is still manual.

---
## Create a Docker Image
Whenever a new tag is created (using the above mentioned steps), a github workflow automatically kicks-in, builds the docker images and pushes to [Dockerhub](https://hub.docker.com/r/nanda/hello-universe/tags?page=1&ordering=last_updated).
