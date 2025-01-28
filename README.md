# livestream-exporter

A framerate and bitrate metrics exporter on Cloud Platform.

## Install

clone src
```bash
git clone xxx.com/livestream-exporter.git
cd livestream-exporter

```

### debug run
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
cp ./config/config.yaml.example ./config/config.yaml

# build and run
make
kubectl create configmap livestream-exporter --from-file=./config/config.yaml
kubectl create deployment livestream-exporter --image=${REGISTRY_DOMAIN}/${REGISTRY_PROJECT}/livestream-exporter:${APP_VERSION} -r 1
```
