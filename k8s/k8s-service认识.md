## k8s service

### 定义
- 定义一组pod 的访问规则，我们通过Controller 创建pod ,通过Service 创建规则取访问pod

### 意义
因为Pod 升级 扩容中，会发生ip 变更，使用service 能够访问到pod
