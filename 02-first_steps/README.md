Tasks:
    * Deploy the pod onto the minikube
    * Find pod's IP address, hosting node, containers, labels, resource usage.
    * Open dashboard `minikube dashboard`
    * Look at the pod logs
    * Record network trafic from within the pod:
        * Hint: kubectl exec, kubectl cp and using tcpdump within the pod. Only the pod's network trafic shall be recorded

Resources:
    * https://kubernetes.io/docs/concepts/workloads/pods/pod/