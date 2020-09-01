# collector

Collects work requests and queues them to be executed

## build and Run

```bash
go build
./collector
```

## docker build

```bash
docker build -t ${GCP_GCR_HOSTNAME}/${GCP_PROJECT_ID}/collector:{VERSION} -f Dockerfile . --build-arg user=${user} --build-arg personal_access_token=${personal_access_token}
```

## docker push

```bash
docker push ${GCP_GCR_HOSTNAME}/${GCP_PROJECT_ID}/collector:{VERSION}
```
