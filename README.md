# GOCATS

This is a sample app to verify and test deployments pipelines in kubernetes

The app is a simple api without any frontend. It exposes the following endpoints

* / \
 Main interface, will reply with a random cat name on GET requests 
* /hello \
 This endpoint replies to GETS with a configurable string, set by the APP_MSG env variable for the app 
* /metrics \
Prometheus metrics for the go process, including a gocats_processed_cats_total metric to see the number of served catnames 

### Configuration

You can configure the application port, the configurable message and the log level using ENV vars: \

Image Defaults:
* APP_PORT=8000
* APP_MSG="Howdy wrangler" 
* LOG_LEVEL=MUTE

By changing LOG_LEVEL to "CATS", the app will log what cat names have been served, and to which client

## Testing 

### Running locally

See the included compose.yml file for configuration. To get GOCATS running, simply

```
podman-compose up (or docker-compose)
```


### Running in Kubernetes 

#### Basic

To run a simple deployment using kubectl, simply apply the manifests in 'example-manifests' directory. This creates a Deployment, a ConfigMap to configure the app and a service to expose the endpoints 

```
kubectl apply -f example-manifest/*.yml -n <your namespace>
```
#### Helm
The directory 'gocats-chart' contains a simple helm chart to deploy the app. Using value overrides, you can modify the config stanza before installing the chart

```yaml
config:
  message: "This is the default app message, served at "
  port: "9000"
  log_level: "CATS"

```
```
helm install gocats ./gocats-chart -n <your namespace>
```
