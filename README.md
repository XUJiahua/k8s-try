
实验

## 使用 minikube 创建 kubernetes 集群（单节点）

```
# 创建 VM，IP 地址 192.168.33.100
vagrant up

# 登录 VM
vagrant ssh

# 以下为VM终端
cd /vagrant

# 安装环境依赖，minikube, kubectl, docker
make env-setup

# 改成 root 用户执行后续命令（方便试验）
sudo su 

# 使用 minikube 创建 k8s 集群（跳过 etcd 等组件的安装）
make minikube-start
```

### 打开 dashboard

dashboard的作用：

1. 资源可视化
2. 在线伸缩Pod
3. 在线提交Resource Config (yaml)或提交表单的形式部署应用

```
minikube dashboard

🤔  Verifying dashboard health ...
🚀  Launching proxy ...
🤔  Verifying proxy health ...
http://127.0.0.1:42837/api/v1/namespaces/kubernetes-dashboard/services/http:kubernetes-dashboard:/proxy/
```

宿主机无法访问VM内的lo网卡。为了方便使用dashboard，操作如下。

```
kubectl proxy --address='0.0.0.0' --disable-filter=true
```

http://192.168.33.100:8001/api/v1/namespaces/kubernetes-dashboard/services/http:kubernetes-dashboard:/proxy/

Ref: https://stackoverflow.com/questions/47173463/how-to-access-local-kubernetes-minikube-dashboard-remotely

## 使用 Docker 进行 Golang 应用的容器化，示例 echoserver

```
cd echoserver

# 打包成Docker image
make build

# 容器运行应用
make start

# 访问服务
curl localhost:7893
```

## 使用 kubectl （apply yaml） 部署（包括更新） echoserver 到集群，体验应用伸缩

```
# 创建n份应用与Service （CRUD之CU）
kubectl apply -f echoserver/echoserver.yaml

# CRUD之R get list
# kubectl get deployments
# kubectl get services
kubectl get pods

NAME                                 READY   STATUS    RESTARTS   AGE
echoserver-54977dcf58-dskws          1/1     Running   0          43h
echoserver-54977dcf58-lqtg5          1/1     Running   0          43h
echoserver-54977dcf58-t7tb9          1/1     Running   0          24h

# CRUD之R get detailed
kubectl describe pod echoserver-54977dcf58-dskws

Name:         echoserver-54977dcf58-dskws
Namespace:    default
Priority:     0
Node:         ubuntu-bionic/10.0.2.15
Start Time:   Wed, 12 Aug 2020 07:52:24 +0000
Labels:       app=echoserver
              pod-template-hash=54977dcf58
              tier=backend
Annotations:  <none>
Status:       Running
IP:           172.17.0.8

...

```

### 自动恢复实例数

echoserver 的实现带有缺陷，/kill 接口会触发panic。
即便如此，k8s检查到pod数不满足配置要求，会自动恢复。

```
# 获取service地址
kubectl get svc

NAME                TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
echoserver          NodePort    10.111.55.193    <none>        7893:31126/TCP   47h

# 正常请求
curl localhost:31126

# 自杀请求
curl localhost:31126/kill

# 查看 Pod 状态
kubectl get pods

NAME                                 READY   STATUS    RESTARTS   AGE
echoserver-54977dcf58-2nhxf          1/1     Running   0          4m59s
echoserver-54977dcf58-4h4wl          0/1     Error     1          4m59s
echoserver-54977dcf58-kr4cc          1/1     Running   0          4m59s

# 自动恢复
kubectl get pods

NAME                                 READY   STATUS    RESTARTS   AGE
echoserver-54977dcf58-2nhxf          1/1     Running   0          5m22s
echoserver-54977dcf58-4h4wl          1/1     Running   2          5m22s
echoserver-54977dcf58-kr4cc          1/1     Running   0          5m22s
```


### Service的类型
K8S VM 内可以访问 ClusterIP，包括 Service的和Pod的。
非K8S集群的主机是无法访问 ClusterIP的。需要通过NodePort/LoadBalancer的形式暴露服务。

NodePort 会在cluster内的每个node打开端口，访问任一node_ip+node_port就可访问service。但是不建议生产使用。
The NodePort abstraction is intended to be a building block for higher-level ingress models (e.g., load balancers). It is handy for development purposes, however, when you don’t need a production URL.

LoadBalancer : exact implementation of a LoadBalancer is dependent on your cloud provider

## demo-consul-101 

使用 Service 来服务发现的例子。

见 demo-consul-101/README.md


