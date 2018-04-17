Tasks:
    * Deploy the pod onto the minikube
    * Find pod's IP address, hosting node, containers, labels, resource usage.
        * `kubectl get pods -o wide`
        * `kubectl describe pod`
    * Open dashboard `minikube dashboard`
    * Look at the pod logs
        * `kubectl logs -f <pod name>`
        * What does `-p` flag does? Why is it useful?
    * Connect to the pod
        * `kubectl port-forward <pod name> local_port:pod_port`
        * Run `curl localhost:local_port`. What do you get?
    * Record network trafic from within the pod:
        * Hint: kubectl exec, kubectl cp and using tcpdump within the pod. Only the pod's network trafic shall be recorded

Resources:
    * https://kubernetes.io/docs/concepts/workloads/pods/pod/
    * https://github.com/kubernetes/dashboard [For deploying dashboard outside minikube]
