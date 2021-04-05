# website operator

Rewrite [luksa/k8s-website-controller](https://github.com/luksa/k8s-website-controller) in `k8s in action`.

## version

- minikube: 1.20.0
- k8s api: 0.20.0
- golang: 1.15

## docker

You can build and push image by yourself

```shell script
docker build -f docker/Dockerfile -t gaoxinge/website-controller.
docker push gaoxinge/website-controller:latest
```

or directly use

- [docker hub](https://hub.docker.com/r/gaoxinge/website-controller)

## yaml

```shell script
# base on https://github.com/luksa/kubernetes-in-action/tree/master/Chapter18
kubectl create -f yaml/website-crd.yaml
kubectl create -f yaml/website-controller.yaml
kubectl create serviceaccount website-controller
kubectl create clusterrolebinding website-controller --clusterrole=cluster-admin --serviceaccount=default:website-controller
kubectl create -f yaml/website-example.yaml
kubeclt delete website kubia
```

## test

```shell script
# http://127.0.0.1:58599 is kubia svc host
curl http://127.0.0.1:58599
curl http://127.0.0.1:58599/subdir
```

## github

- [github repository example](https://github.com/luksa/kubia-website-example)
- [github proxy](https://blog.csdn.net/weixin_42886104/article/details/106454331)

## TODO

- [ ] add polling to website control for event missing
- [ ] add update to website control