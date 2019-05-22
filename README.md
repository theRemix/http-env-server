# http-env-server

Serves the content of env vars for testing rolling deployments

## Run Locally

```sh
env PORT=3000 VERSION=1.0.0 go run main.go
```

## Run with docker

```sh
docker run -d \
  -p 3000:80 \
  -e VERSION=1.0.0 \
  theremix/http-env-server
```
