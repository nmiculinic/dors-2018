Install minikube, kubectl, go (most examples use go here)
* https://kubernetes.io/docs/tasks/tools/install-minikube/
* https://kubernetes.io/docs/tasks/tools/install-kubectl/
* https://golang.org/doc/install#install

```bash
minikube start --vm-driver=virtualbox
kubectl config get-contexts  # Check active context
eval $(minikube docker-env)  # set minikubes docker daemon as currently running (ala docker-machine)
docker ps 
kubectl get pod --all-namespaces
kubectl cluster-info
kubectl get events
kubectl config view
minikube addons enable heapster
minikube addons open heapster
```

More information:
* https://kubernetes.io/docs/tutorials/kubernetes-basics
* https://kubernetes.io/docs/getting-started-guides/minikube/
* https://kubernetes.io/docs/reference/kubectl/docker-cli-to-kubectl/
* https://kubernetes.io/docs/user-guide/walkthrough/
* https://kubernetes.io/docs/user-guide/walkthrough/k8s201/
* https://github.com/GoogleCloudPlatform/kubernetes-workshops (( Some of the materials were inspired by it ))