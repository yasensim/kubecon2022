# kubecon2022
Demo App for Kubecon EU 2022 for Istio Service Mesh and Argo Rollout


[Demo link](https://youtu.be/5Ko-CnP2qhA?list=PLj6h78yzYM2MCEgkd8zH0vJWF7jdQ-GRR&t=1318)

## validations windows
```
while true ; do curl -w " - code: %{http_code}\n" http://kc.yasensim.net/kubecon ; sleep 0.5 ; done
```

## commands used 

preinstalled applications  
- argo rollout 
- istio 

```
kubectl argo rellout version
kubectl create namespace kubecon
kubectl label  namespace kubecon istio-injection=enabled 
helm upgrade --install kubecon helm --namespace helm --values rollout/values-analysis.yaml --set ingress.host=kc.yasensim.net --set image.tag=v1.0 --wait
```


