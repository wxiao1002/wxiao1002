# 操作集群常见命令

master 查看集群节点数 ：kubectl get nodes

master查看集群所有组件的情况：kubectl get pods -n kube-system -o wide

master 删除集群某个节点 ：kubectl delete nodes nodeName

node节点呗抛弃出集群，重置节点：kubeadm reset

如果是想重新搭建，master记得删除 $HOME/.kube 目录及其其下的所有数据
master欢迎新的节点，抛出橄榄枝：kubeadm token create --print-join-command

关于集群中正确删除Pod的操作：先删除Pod ，再删除相应的Deployment
kubectl get pods  　　          查看当前的pod
　　kubectl delete pod podName
kubectl get deployments      查看当前部署的应用
　　kubectl delete depliyment depName

## 创建一个pod 并验证是否成功

kubectl create deployment nginx --image=nginx

kubectl expose deployment nginx --port=80 --type=NodePort

kubectl get pod,svc

然后访问任意节点ip + 端口就能访问nginx
