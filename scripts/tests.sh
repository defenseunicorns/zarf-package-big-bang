#!/bin/bash
#

kubectl cluster-info
echo "current-context:" $(kubectl config current-context)

kubectl get nodes

kubectl wait kustomization/bigbang -n flux-system --for=condition=Ready --timeout=1200s
