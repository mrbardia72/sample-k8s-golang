# sample-k8s-golang

### 1 - build image
#### docker images build -t kuber-bardia .

### 2 - login & push to docker hub
#### docker login
#### docker tag kuber-bardia USERNAME/kuber-bardia
#### docker push USERNAME/kuber-bardia

### 3 - create service & deployment
#### kubectl create -f service.yml
```
service/service-kuber-bardia created
```
#### kubectl create -f deployment.yaml
``` 
deployment.apps/deploy-kuber-bardia created
```
### 4 - get information
#### kubectl get pods -o wide
```
NAME                                   READY   STATUS    RESTARTS   AGE   IP            NODE       NOMINATED NODE   READINESS GATES
deploy-kuber-bardia-6c7b4cbfff-7phlp   1/1     Running   0          23h   10.244.0.65   minikube   <none>           <none>
deploy-kuber-bardia-6c7b4cbfff-9hhd2   1/1     Running   0          23h   10.244.0.73   minikube   <none>           <none>
deploy-kuber-bardia-6c7b4cbfff-g8q52   1/1     Running   0          23h   10.244.0.69   minikube   <none>           <none>
deploy-kuber-bardia-6c7b4cbfff-mlsh2   1/1     Running   0          23h   10.244.0.70   minikube   <none>           <none>
deploy-kuber-bardia-6c7b4cbfff-v89mj   1/1     Running   0          23h   10.244.0.66   minikube   <none>           <none>
deploy-kuber-bardia-6c7b4cbfff-wbnn2   1/1     Running   0          23h   10.244.0.72   minikube   <none>           <none>

```
#### kubectl get services -o wide
```
service-kuber-bardia   NodePort    10.104.215.15   <none>        80:30009/TCP   9m59s   app=test-service-kuber-bardia
```
#### kubectl get deployments -o wide
```
deploy-kuber-bardia   6/6     6            6           23h   container-bardia   xxgrg52mshs8945q/kuber-bardia   type=back-end
```

### 5 - for delete service & deployment
#### kubectl delete deployment deploy-kuber-bardia
``` 

```
#### kubectl delete service deploy-kuber-bardia
``` 

```