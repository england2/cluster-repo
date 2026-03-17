# initialize argocd

kubectl apply --server-side -k 'https://github.com/argoproj/argo-cd/manifests/crds?ref=stable'
