# zarf-package-big-bang
Pre-built Zarf Package of Big Bang core. The current version of Big Bang supported is 1.43.0

## Prerequisites
- Zarf is installed. Current version used is: [v0.21.3](https://github.com/defenseunicorns/zarf/releases/tag/v0.21.3)
- Optional: A working Kubernetes cluster. e.g KinD, k3d, k3s, etc. If necessary, Zarf can be used to deploy a built-in k3s distribution.

## Build the package

```
$ make build
...
$ ls -l build/
total 2499764
-rw-r--r--. 1 user user 2559751131 Oct 12 08:24 zarf-package-big-bang-amd64.tar.zst
$ 
```

## Deploy the package
1. Download and deploy Zarf's init package.
```
$ zarf package deploy ~/Downloads/zarf-init-amd64.tar.zst --components git-server --confirm
```

Optionally, if a Kubernetes cluster is not installed, we can tell Zarf to install one for us:

```
zarf package deploy ~/Downloads/zarf-init-amd64.tar.zst --components k3s,git-server --confirm
```

2. Deploy the Big Bang package created in the Build step above:
```
zarf package deploy ./build/zarf-package-big-bang-amd64.tar.zst --confirm
```

Check whether the deployement succeeded. If the deployment is successful, then you should see this message from the HelmRelease resource in the cluster.

```
$ kubectl get helmrelease -A
bigbang     bigbang           72m   True    Release reconciliation succeeded
bigbang     cluster-auditor   71m   True    Release reconciliation succeeded
bigbang     gatekeeper        71m   True    Release reconciliation succeeded
bigbang     istio             71m   True    Release reconciliation succeeded
bigbang     istio-operator    71m   True    Release reconciliation succeeded
bigbang     jaeger            71m   True    Release reconciliation succeeded
bigbang     keycloak          71m   True    Release reconciliation succeeded
bigbang     kiali             71m   True    Release reconciliation succeeded
bigbang     loki              71m   True    Release reconciliation succeeded
bigbang     metrics-server    71m   True    Release reconciliation succeeded
bigbang     minio             71m   True    Release reconciliation succeeded
bigbang     minio-operator    71m   True    Release reconciliation succeeded
bigbang     monitoring        71m   True    Release reconciliation succeeded
bigbang     promtail          71m   True    Release reconciliation succeeded
```

## Day 2.
### Enable SSO

Navigate to the Day 2 sections of the `kustomizations/values.yaml` file and fill in the details in the `sso` and `authservice` sections to
enable SSO for your application.
