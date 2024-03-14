### Create your connection configuration and secrets
##### Prepare the Secret Data: For PostgreSQL, this usually involves the database user password. Encode your password using Base64.
##### For example: echo -n 'password' | base64
```shell
 kubectl apply -f postgres-secrets.yml
 kubectl get secret postgres-secret-config -o yaml
```

### Create PersistentVolume
```shell
❯ kubectl apply -f pv-volume.yml
❯ kubectl get pv postgres-pv-volume
```

### Create PersistentVolumeClaim
```shell
❯ kubectl apply -f pv-claim.yml
❯ kubectl get pvc postgres-pv-claim
```

### Create Kubernetes deployment
```shell
❯ kubectl apply -f postgres-deployment.yml
❯ kubectl get deployments
```

### Create service
```shell
kubectl apply -f postgres-service.yml
kubectl get service postgres
```

### Test the connection to the Postgres database
```shell
❯ kubectl get pods
NAME READY STATUS RESTARTS AGE
postgres-57f47 1/1 Running 0 30m

❯ kubectl exec -it postgres-57f47 -- psql -U postgres
```
### test set configmap
```shell
show max_connections
```
### resource
https://www.sumologic.com/blog/kubernetes-deploy-postgres
https://refine.dev/blog/postgres-on-kubernetes