# livestream-exporter

A framerate and bitrate metrics exporter on Cloud Platform.

## Install

clone src
```bash
git clone xxx.com/livestream-exporter.git
cd livestream-exporter
```

### debug

```bash
go run main.go
```

### build and run

```bash
# set environment
export APP_VERSION=1.0
export APP_PORT=8080
export REGISTRY_DOMAIN=domain_var
export REGISTRY_PROJECT=project_var

# build image
make

# run on k8s
cp ./config/config.yaml.example ./config/config.yaml
kubectl create configmap livestream-exporter --from-file=./config/config.yaml
kubectl create deployment livestream-exporter --image=${REGISTRY_DOMAIN}/${REGISTRY_PROJECT}/livestream-exporter:${APP_VERSION} -r 1
# patch configmap and volume
kubectl patch deployments livestream-exporter -p '{"spec": {"template": {"spec": {"imagePullSecrets": [{"name": "harbor-secret"}]}}}}'
kubectl patch deployment livestream-exporter --type='json' -p='[{"op": "add", "path": "/spec/template/spec/volumes", "value": [{"name": "livestream-exporter", "configMap": {"name": "livestream-exporter", "items": [{"key": "config.yaml", "path": "config.yaml"}]}}]}, {"op": "add", "path": "/spec/template/spec/containers/0/volumeMounts", "value": [{"name": "livestream-exporter", "mountPath": "/app/config/config.yaml", "subPath": "config.yaml"}]}]'
```
