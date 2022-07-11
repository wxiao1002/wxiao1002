# Trie 树
也称为字典树。它是一个树形结构，用来解决在一组字符串集合中快速查找某个字符串的问题

## 本质
利用字符串的公共前缀，将重复的前缀组合起来

## 代码实现
如何存储子节点，我们使用经典的存储方式使用 a-z 用下标存储每个字符开始的数据
```java
package Test;

class TrieNode {
    char data;
    
     TrieNode [] children  = new TrieNode[26];

    public TrieNode (char data){
        this.data = data;
    }
    
}

public class TrieTree {
    
    private TrieNode root = new TrieNode('/');
    
    
    public void  add (char [] text){
        TrieNode p = root;
        for (cha temp : text){
            int index = temp - 'a';
            if (p.children[index] == null){
                p.children[index] = new TrieNode(temp);
            }
            p = p.children[index];
        }
    }
    
    public boolean find(char [] text){
        TrieNode p = root;
        for (cha temp : text){
            int index = temp - 'a';
            if (p.children[index] == null){
                return false;
            }
            p = p.children[index];
        }
        // 判断p 是否是叶子
        return !p.isLeaf();
    }
}
```
