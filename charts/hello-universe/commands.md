## Helm Commands
* [Lint Helm templates](#lint-helm-templates)
* [Render Helm templates](#render-helm-templates)
* [Do a Helm Release](#do-a-helm-release)

---
### Lint Helm templates
```sh
cd charts/hello-universe
helm lint -v 5 \
    --logtostderr \
    --debug \
    --namespace hello-universe \
    --strict \
    --values values.yaml \
    --values values-secrets.yaml \
    --values values-infra-secrets-dev.yaml \
    --values env-overrides/values-dev.yaml \
    .
```

---
### Render Helm templates
```sh
cd charts/hello-universe
helm template -v 5 \
    --namespace hello-universe \
    --logtostderr \
    --debug \
    --show-only templates/deployment.yaml \
    --values values.yaml \
    --values values-secrets.yaml \
    --values values-infra-secrets-dev.yaml \
    --values env-overrides/values-dev.yaml \
    hello-universe \
    .
```
> Drop `show-only` flag mentioned above, to render all templates in one go.

---
### Do a Helm release
```sh
cd charts/hello-universe
helm upgrade -v 3 \
    --namespace hello-universe \
    --logtostderr \
    --debug \
    --install \
    --atomic \
    --timeout 60s \
    --debug \
    --dry-run \
    --cleanup-on-fail \
    --values values.yaml \
    --values values-secrets.yaml \
    --values values-infra-secrets-dev.yaml \
    --values env-overrides/values-dev.yaml
    v1 \
    .
```
> Drop `dry-run` flag mentioned above, to do an actual release.
