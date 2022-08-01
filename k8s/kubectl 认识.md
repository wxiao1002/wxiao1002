# kubectl

## 概述
kubectl 是k8s 集群中的命令航工具，通过kubectl 能够实现对集群本身进行管理和维护

## kubectl 命令的语法

大致语法内容是是： kubectl [command] [type] [name] [flags]

- command: 操作指令，比如 create,get,describe,delete
- type： 资源类型，大小写敏感，但可以是单数、复数、或者缩略形式

    - kubectl get pod nginx
    - kubectl get pods nginx
    - kubectl get po nginx
- name: 指定资源的名称，大小写敏感，可省略，则删除所有的资源
- flags: 指定可选的参数
    - 比如可用 -s 或者-server 参数指定 k8s api server 的地址和端口

## kubectl --help

- 输入 kubectl --help 可输入如下内容

```

Basic Commands (Beginner):
  create        Create a resource from a file or from stdin.
  expose        使用 replication controller, service, deployment 或者 pod 并暴露它作为一个新的 Kubernetes
Service
  run           在集群中运行一个指定的镜像
  set           为 objects 设置一个指定的特征
      
Basic Commands (Intermediate):
  explain       查看资源的文档
  get           显示一个或更多 resources
  edit          在服务器上编辑一个资源
  delete        Delete resources by filenames, stdin, resources and names, or by resources and label selector
      
Deploy Commands:
  rollout       Manage the rollout of a resource
  scale         Set a new size for a Deployment, ReplicaSet or Replication Controller
  autoscale     自动调整一个 Deployment, ReplicaSet, 或者 ReplicationController 的副本数量
      
Cluster Management Commands:
  certificate   修改 certificate 资源.
  cluster-info  显示集群信息
  top           Display Resource (CPU/Memory/Storage) usage.
  cordon        标记 node 为 unschedulable
  uncordon      标记 node 为 schedulable
  drain         Drain node in preparation for maintenance
  taint         更新一个或者多个 node 上的 taints
      
Troubleshooting and Debugging Commands:
  describe      显示一个指定 resource 或者 group 的 resources 详情
  logs          输出容器在 pod 中的日志
  attach        Attach 到一个运行中的 container
  exec          在一个 container 中执行一个命令
  port-forward  Forward one or more local ports to a pod
  proxy         运行一个 proxy 到 Kubernetes API server
  cp            复制 files 和 directories 到 containers 和从容器中复制 files 和 directories.
  auth          Inspect authorization
      
Advanced Commands:
  diff          Diff live version against would-be applied version
  apply         通过文件名或标准输入流(stdin)对资源进行配置
  patch         使用 strategic merge patch 更新一个资源的 field(s)
  replace       通过 filename 或者 stdin替换一个资源
  wait          Experimental: Wait for a specific condition on one or many resources.
  convert       在不同的 API versions 转换配置文件
  kustomize     Build a kustomization target from a directory or a remote url.
      
Settings Commands:
  label         更新在这个资源上的 labels
  annotate      更新一个资源的注解
  completion    Output shell completion code for the specified shell (bash or zsh)
      
Other Commands:
  alpha         Commands for features in alpha
  api-resources Print the supported API resources on the server
  api-versions  Print the supported API versions on the server, in the form of "group/version"
  config        修改 kubeconfig 文件
  plugin        Provides utilities for interacting with plugins.
  version       输出 client 和 server 的版本信息
      
Usage:
  kubectl [flags] [options]
```

## kubectl 命令总结分类

| 命令      | 解释                      |
|---------|-------------------------|
 | create  | 通过文件或者标准输入创建资源          | 
 | expose  | 将资源公开为一个新的Service       |
| run     | 在集群中运行一个特定的对象           |
| get     | 显示一个或者多个资源              |
| set     | 在对象上设置特别的功能             |
| explain | 文档参考资料                  |
| edit    | 使用默认编辑器编辑一个资源           |
| delete  | 通过文件名、标准输入，资源名称或者标签删除资源 |

## 部署命令

| 命令             | 解释                                        |
|----------------|-------------------------------------------|
| rollout        | 管理资源的发布                                   | 
| rolling-update | 对给定的复制控制器滚动更新                             |
| scale          | 扩容或者缩容pod 数量，Deployment\ReplicaSet,RC,Job |
| autoscale      | 创建一个自动惊醒扩容和缩容并设置Pod 数量                    |
       

## 集群管理命令

| 命令           | 解释                   |
|--------------|----------------------|
| certificate  | 修改证书资源               | 
| cluster-info | 显示集群信息               |
| top          | 显示资源占用，需要Heapster 运行 |
| cordon       | 标记节点不可被调度            |
| uncordon     | 标记节点可被调度             |
| drain        | 驱逐节点上的应用，准备下线维护      |
| taint        | 修改节点 taint标记         |
      