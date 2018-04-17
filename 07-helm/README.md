Introduction:
    * Why helm -- The Kubernetes Package Manager?
        * Sooner or later, your app shall depend on something (database, cache layer, etc.)
        * With helm you can package those dependencies within the helm chart
        * Plus templating features for common k8s .yaml's
        * It uses go template engine by default with
Tasks:
    * Install helm --  https://docs.helm.sh/using_helm/
    * helm init -- to install Tiller
    * Look over provided charts: https://github.com/kubernetes/charts
    * Create new helm chart `helm create <chart name>`
    * Set templates for image and replicas count (( of the app we've been working on ))
    * Upgrade a chart `helm upgrade` (( increase the number of replicas ))
    * Look at chart deployment history `helm history`
    * Get values `helm get values`. See other options `--help`, notably --revision
    * Rollback a chart `helm rollback`
Resources:
    * https://helm.sh/
