credits: https://github.com/hashicorp/demo-consul-101

## Non-Docker

Start counting service.

```
PORT=9001 go run main.go
```

Start dashboard service.

```
PORT=8080 go run main.go
```


## Docker (Compose)

To run both microservices with Docker Compose.

```
$ cd demo-consul-101
$ docker-compose up
```

You can view the operational application dashboard at http://localhost:8080

## Kubernetes

```
kubectl apply -f counting-service.yaml

service/counting-service created
deployment.apps/counting-service created
```


```
kubectl apply -f dashboard-service.yaml

service/dashboard-service created
deployment.apps/dashboard-service created
```


```
kubectl get svc

NAME                TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
counting-service    ClusterIP   10.98.131.201    <none>        9001/TCP         111s
dashboard-service   NodePort    10.106.75.137    <none>        8080:30267/TCP   103s
```

You can view by visiting: http://node_ip:30267
