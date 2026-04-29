# misc dev notes

# install argocd (https://argo-cd.readthedocs.io/en/stable/getting_started/)

kubectl create namespace argocd
kubectl apply -n argocd --server-side --force-conflicts -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# get kube config
mkdir -p ~/.kube
sudo cp /etc/rancher/k3s/k3s.yaml ~/.kube/config
sudo chown (id -u):(id -g) ~/.kube/config
chmod 600 ~/.kube/config

# portforward argocd ui

kubectl port-forward svc/argocd-server -n argocd 8080:443

# get default argocd ui password

kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo

# sealed secrets
