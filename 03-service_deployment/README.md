Tasks:
    * find service endpoints `k get endpoints hello-world`
    * connect to the pod inside the cluster (( `kubectl run -it --rm --image=alpine test` ))
    * Try using headless service (clusterIP:None), clusterIP and NodePort ones. What is the difference between those?
    How many IP does `dig +short +search` returns?
    * Using NodePort (( `minikube ip`, `kubectl get svc` )) connect to the service.
    (find nodeport and connect to any k8s node, e.g. minikube's one)
    * Find service endpoints `k get endpoints`

Resources:
    https://kubernetes.io/docs/concepts/workloads/controllers/deployment/
    https://kubernetes.io/docs/concepts/services-networking/service/
    * Using dig (part of bind-tools) -- https://www.madboa.com/geek/dig/#use-the-search-list-in-etc-resolv-conf