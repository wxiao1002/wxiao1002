## k8s  namespace

### namespace 概述
- 主要实现多用户资源隔离，形成逻辑上的分组，便于不同分组在共享整个集群的资源时便于管理
- k8s 在集群启动后，会创建一个default 的Namespace,如果不特别指出namespace 则用户创建的Pod
Rc，Service 都会创建到这个默认的名称为default 的namespace中

### namespace 创建
- 创建一个名称为 testnamespace 名称的命名空间
kubectl create namespace testnamespace
- 查看所有的namespace
kubectl get namespace 

### namespace 使用
比如我们创建一个pod ，只需要在气候加一个属性即可实现Pod 分类
- kubectl apply -f test.yaml -n testnamespace

在名称为testnamespace namespace 下通过test.yaml 部署

- kubectl get pods -n testnamespace

查看namespace 名称下所有的pod