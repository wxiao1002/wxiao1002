![](https://img2020.cnblogs.com/blog/1526344/202101/1526344-20210118143942091-1641998139.png)

语法格式
通过缩进表示层级关系

不能使用Tab进行缩进，只能使用空格

一般开开头使用两个空格

字符key后缩进一个空格 key： value

使用 --- 表示一个新的yaml文件的开始


![](https://img2020.cnblogs.com/blog/1526344/202101/1526344-20210127181041619-929493713.png)

## 使用yaml 部署应用

kubectl apply -f web1.yaml


## 第三部对外发布(暴露端口)

kubectl expose deployment web1 --port=80 --type=NodePort --target-port=80 --name=web1 -o yaml > web1Expose.yaml