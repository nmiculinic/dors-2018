Tasks:
    * find service endpoints `k get endpoints hello-world`
    * connect to the pod inside the cluster (( `kubectl run -it --rm --image=alpine test` ))
    * Use NodePort (( `minikube ip`, `kubectl get svc` ))
Resources:
    https://kubernetes.io/docs/concepts/workloads/controllers/deployment/
    https://kubernetes.io/docs/concepts/services-networking/service/