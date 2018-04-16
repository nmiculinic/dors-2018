Tasks:
    * `kubectl run -it --rm --command --image-pull-policy=Never --image=dors-hello-world test bash`
    * Try deleting a pod,and see what happens. What was the IP of the pod before? What's the IP of the pod now?
    * kubectl get pvc, pv
    * After clearing and deleting all resources `make clean` what happens to the data? Run `kubectl get pv` and notice something
    * Create a config map and replace hardcoded DORS_HOME envorinment variable
Resources:
    * https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/
    * https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/