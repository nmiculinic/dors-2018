Tasks:
    * find service endpoints `k get endpoints hello-world`
    * Create a new pod and open bash within it
        e.g. (( `kubectl run -it --rm --command --image-pull-policy=Never --image=dors-hello-world test bash ))
    * What's the difference between following service types:
        How many IP does `dig +short +search` returns in each case?
        * ClusterIP
        * NodePort

        Optionally with following 3 further types:
        * Headless service -- no k8s load balancing nor static IP address
        * LoadBalancer
        * ExternalName
    * Using NodePort (( `minikube ip`, `kubectl get svc` )) connect to the service.
    (find nodeport and connect to any k8s node, e.g. minikube's one)
    * Find service endpoints `kubectl get endpoints`

Resources:
    https://kubernetes.io/docs/concepts/workloads/controllers/deployment/
    https://kubernetes.io/docs/concepts/services-networking/service/
    * Using dig (part of bind-tools) -- https://www.madboa.com/geek/dig/#use-the-search-list-in-etc-resolv-conf