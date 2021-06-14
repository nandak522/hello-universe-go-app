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
    --namespace webhooks \
    --logtostderr \
    --debug \
    --values values.yaml \
    hello-universe \
    .

helm upgrade -v 3 \
    --namespace hello-universe \
    --logtostderr \
    --debug \
    --install \
    --atomic \
    --timeout 60s \
    --debug \
    --cleanup-on-fail \
    --values values.yaml \
    hello-universe \
    .
```
