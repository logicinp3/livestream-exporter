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
mv ./config/config.yaml.example ./config/config.yaml

# run in src
go run main.go

# run in k8s
kubectl create secret generic livestream-exporter --from-file=./config/config.yaml
kubectl create deployment livestream-exporter --image=${REGISTRY_DOMAIN}/${REGISTRY_PROJECT}/livestream-exporter:${APP_VERSION} --replicas=1
kubectl patch deployment livestream-exporter -p '{"spec": {"template": {"spec": {"imagePullSecrets": [{"name": "harbor-secret"}]}}}}'
kubectl patch deployment livestream-exporter --type='json' -p='[{"op": "add", "path": "/spec/template/spec/volumes", "value": [{"name": "config-volume", "secret": {"secretName": "livestream-exporter", "items": [{"key": "config.yaml", "path": "config.yaml"}]}}]}, {"op": "add", "path": "/spec/template/spec/containers/0/volumeMounts", "value": [{"name": "config-volume", "mountPath": "/app/config/config.yaml", "subPath": "config.yaml"}]}]'
```

### deploy

```bash
# set environment
export APP_VERSION=1.0
export APP_PORT=8080
export REGISTRY_DOMAIN=domain_var
export REGISTRY_PROJECT=project_var

# build image
make

# run in k8s
sed -i "s#image:.*#image: $REGISTRY_DOMAIN/$REGISTRY_PROJECT/livestream-exporter:$APP_VERSION#" ./kustomize/deployment.yaml
mv ./config/config.yaml.example ./config/config.yaml && kubectl create secret generic livestream-exporter --from-file=./config/config.yaml
kubectl apply -k ./kustomize/
```
