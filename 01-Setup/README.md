Install minikube, kubectl
```bash
minikube start --vm-driver=virtualbox
kubectl config get-contexts  # Check active context
eval $(minikube docker-env)  # set minikubes docker daemon as currently running (ala docker-machine)
docker ps 
kubectl get pod --all-namespaces
```

More information:
https://kubernetes.io/docs/getting-started-guides/minikube/
