## k8s service

### 定义
- 定义一组pod 的访问规则，我们通过Controller 创建pod ,通过Service 创建规则取访问pod

### 意义
1. 因为Pod 升级 扩容中，会发生ip 变更，使用service 能够访问到pod
2. 定义一组pod 访问策略（负载均衡）
service 会对一组Pod做负载均衡

### Pod 与Service 的关系

- Controller 与Pod 是通过标签label和选择器Selector建立关联
- service 与Controller 如出一撤与Pod 建立关系

![](https://img2020.cnblogs.com/blog/1526344/202101/1526344-20210127142548589-193102498.png)


### 常见的Service类型

1. ClusterIP
- 一般用于集群内部使用，比如前面说的前端pod 访问后端pod,使用的是集群内部的ip,对外无效

2. NodePort
- NodePort 是对外访问应用时候使用，比如我们前端的pod 是要对外暴露的，面向集群外部访问

3. LoadBalancer
- 对外访问时使用，公有云的服务调用