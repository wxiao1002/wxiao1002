## k8s Controller

### 概念
Controller 不像是Pod，是虚拟出来的概念，Controller 是实际存在的，在k8s 在集群上管理和运行容器的对象

### Pod 与Controller

Pod 是通过Controller 实现应用的运维，比如伸缩，滚动升级

####  Pod 和Controller 是如何取的联系
- 通过 label 标签建立联系

https://img2020.cnblogs.com/blog/1526344/202101/1526344-20210127181032248-1933884124.png
