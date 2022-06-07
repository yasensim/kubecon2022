# kubecon2022
Demo App for Kubecon EU 2022 for Istio Service Mesh and Argo Rollout



[Demo link](https://youtu.be/5Ko-CnP2qhA?list=PLj6h78yzYM2MCEgkd8zH0vJWF7jdQ-GRR&t=1318)

## validations windows
```
while true ; do curl -w " - code: %{http_code}\n" http://kc.yasensim.net/kubecon ; sleep 0.5 ; done
```

## deploying the new versions app 

### preparation 
preinstalled applications  
- argo rollout 
- istio 
- lable namespace as below 
```
kubectl argo rellout version
kubectl create namespace kubecon
kubectl label  namespace kubecon istio-injection=enabled 
```

### deploy version 1
```
helm upgrade --install kubecon helm --namespace kubecon --values rollout/values-analysis.yaml \
--set ingress.host=kc.yasensim.net --set image.tag=v1.0 --wait

kubectl argo rollouts --namespace kubecon get rollout kubecon-kubecon --watch 
```
![image](https://user-images.githubusercontent.com/4955356/172356136-bcdaf56f-b058-4918-8b03-0dd1d65e1048.png)


```
kubectl get rollout -n kubecon kubecon-kubecon -o yaml 
```

### deploy version 2 - bad application  
```
helm upgrade  kubecon helm --namespace kubecon --reuse-values --set image.tag=v2.0 --wait

#check virtual service
kubectl get vs -n kubecon -o yaml 
kubectl argo rollouts --namespace kubecon get rollout kubecon-kubecon --watch 
```

![image](https://user-images.githubusercontent.com/4955356/172357534-21ebd9c6-484f-4474-be90-f1cc9239cb18.png)



### deploy version 3 - good application 
```
helm upgrade  kubecon helm --namespace kubecon --reuse-values --set image.tag=v2.0 --wait
```

