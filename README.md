# sample-k8s-golang

### 1 - build image
```
$ docker images build -t kuber-bardia .
```

### 2 - login & push to docker hub
```
$ docker login
$ docker tag kuber-bardia USERNAME/kuber-bardia
$ docker push USERNAME/kuber-bardia
``` 

### 3 - apply service & deployment
```
$ kubectl apply -f service.yml
$ kubectl apply -f deployment.yaml
```
### 4 - get information pod & service & deployment
```
$ kubectl get pods -o wide

NAME                                   READY   STATUS    RESTARTS   AGE   IP            NODE       NOMINATED NODE   READINESS GATES
deploy-kuber-bardia-6c7b4cbfff-7phlp   1/1     Running   0          23h   10.244.0.65   minikube   <none>           <none>
deploy-kuber-bardia-6c7b4cbfff-9hhd2   1/1     Running   0          23h   10.244.0.73   minikube   <none>           <none>
deploy-kuber-bardia-6c7b4cbfff-g8q52   1/1     Running   0          23h   10.244.0.69   minikube   <none>           <none>
```

```
$ kubectl get services -o wide
service-kuber-bardia   NodePort    10.104.215.15   <none>        80:30009/TCP   9m59s   app=test-service-kuber-bardia
```

```
$ kubectl get deployments -o wide
deploy-kuber-bardia   6/6     6            6           23h   container-bardia   xxgrg52mshs8945q/kuber-bardia   type=back-end
```

#### 5 - You can also stream the Pod logs by typing the following command:
``` 
$ kubectl logs -f deploy-kuber-bardia-bfbdf88c-qfqzf
```
#### 6 - Scale a Kubernetes deployment
You can scale the number of Pods by increasing the number of replicas in the Kubernetes deployment manifest and applying the changes using kubectl.
You can also use kubectl scale command to increase the number of pods:
``` 
$ kubectl scale --replicas=4 deployment/deploy-kuber-bardia
OR
$ kubectl scale --replicas=4 deployment.yml
```

#### 7 - Delete resources (pod & service & deployment)

``` 
$ kubectl delete pod deploy-kuber-bardia-bfbdf88c-qfqzf
```

``` 
$ kubectl delete service service-kuber-bardia
```

``` 
$ kubectl delete deployment deploy-kuber-bardia
```
