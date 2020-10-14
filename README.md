## Inframap as a Service

This service hosts a webserver in order to consume inframap lib and to generate dot graphs from Terraform State

:warning: disclaimer: this is a PoC and should not be used in production :warning:

### Usage

Once ran, you can access the endpoint with curl:

```shell
curl -X POST localhost:8080 -H "Content-Type: application/text" --data @generate/testdata/flexibleengine.json
```

Where the `data` is a Terraform state.

### Run

```
go run main.go
```

it will start a webserver on the port :8080

### Docker

There is a Dockerfile, you can build your docker image:

```
docker build -t gcr.io/<gcp-project-id>/inframap-service .
```

### Deploy

Example of Google Cloud Run deployment

```
gcloud run deploy inframap-service --image gcr.io/<gcp-project-id>/inframap-service --port 8080 --platform managed --allow-unauthenticated
curl -X POST $(gcloud run services list --platform managed --format json | jq -r '.[] | select( .metadata.name == "inframap-service") | .status.url ') -H "Content-Type: application/text" --data @generate/testdata/flexibleengine.json
```
