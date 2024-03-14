 Create your connection configuration and secrets
 kubectl apply -f postgres-secrets.yml
 kubectl get secret postgres-secret-config -o yaml

 Create PersistentVolume and PersistentVolumeClaim
❯ kubectl apply -f pv-volume.yml
❯ kubectl get pv postgres-pv-volume

❯ kubectl apply -f pv-claim.yml
❯ kubectl get pvc postgres-pv-claim

Create Kubernetes deployment
❯ kubectl apply -f postgres-deployment.yml
❯ kubectl get deployments

Create service
kubectl apply -f postgres-service.yml
kubectl get service postgres


Test the connection to the Postgres database
❯ kubectl get pods
NAME READY STATUS RESTARTS AGE
postgres-57f4746d96-7z5q8 1/1 Running 0 30m

❯ kubectl exec -it postgres-57f4746d96-7z5q8 -- psql -U postgres

There is also a handy way to store the pod name in a variable:
POD=`kubectl get pods -l app=postgres -o wide | grep -v NAME | awk '{print $1}'`

https://www.sumologic.com/blog/kubernetes-deploy-postgres/