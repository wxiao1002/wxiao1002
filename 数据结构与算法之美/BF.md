# BF 算法
常用在字符串匹配
- 主串和模式串：假设在字符串a中查找字符串b,则a 是主，b为模式串
## 概念
- BF 算法也称暴力匹配法
- 如果模式串长度为m,主串长度为n,则在主串中有n-m+1个子串，需要对比这些子串是否与模式串匹配
- 实现
```java
package main;

class BF {
    int bf (char []a, char [] b,int n,int m){
        for (int i=0;i<n-m;i++){
            int j= 0;
            for (;j<m;j++){
                if (a[i+j]!= b[j]){
                    break;
                }
            }
            // 相等则说明没有 break 字符串匹配
            if (j==m){
                return i;
            }
        }
    }
}

```