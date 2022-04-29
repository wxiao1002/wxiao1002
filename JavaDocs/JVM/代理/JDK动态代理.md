## 动态代理
区别于代理模式中的区别：代理类是动态生成的
### 分类
1. 基于接口 </br>
    JDK代理
2. 基于类 </br>
   cglib，Java 字节码实现也叫java assist
    
### Proxy
代理类
### InvocationHandler
- 定义：由代理实例调用程序实现的接口
- 方法：
```
Object invoke(Object proxy,Method method,Object [] args) throws Throwable;
// proxy : 代理类
// method: 目标类 需要 invoke 的方法
// args: 目标类方法的参数
```
### 实现

```java
import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;

class MyInvocationHandler<T> implements InvocationHandler {

    private final T targetObj;

    @Override
    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
        System.out.println("before");
        Object result = method.invoke(targetObj, args);
        System.out.println("after");
        return result;
    }

    public MyInvocationHandler(T targetObj) {
        this.targetObj = targetObj;
    }
}

interface Student {
    void doWork();
}

class XiaoStudent implements Student {
    @Override
    public void doWork() {
        System.out.println("我是小学生");
    }
}

class TestProxy {
    public static void main(String[] args) {
        Student student = new XiaoStudent();
        InvocationHandler invocationHandler= new MyInvocationHandler<Student>(student);
        Student renterProxy = (Student)Proxy.newProxyInstance(Student.class.getClassLoader(),new Class<?>[]{Student.class}, invocationHandler);
        renterProxy.doWork();
    }
}

```
### 优化项与思考
被代理类只可以直接调用反射生成，不必new,当然接口是如何工作了，被代理类不能直接new 是如何做到处理？

### Mybatis 的动态代理
#### 动态代理类后结构
1. 结构：生成后的类 继承于Proxy 实现目标类共有的接口，并实现接口方法，使用成员变量记录目标类的方法，接口方法实现中使用 
```
h.inboke(this,method,args [];
```
2. 调用流程
```
proxy.doWork() ----->InvocationHandler.invoke()----->method.invoke()---->target.doWork()
```
#### Mybatis 为啥不需要实现类

```java
import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;

class ProxyTest {
    public static void main(String[] args) {
        Student o = (Student)Proxy.newProxyInstance(Student.class.getClassLoader(), new Class[]{Student.class}, new MapperInvocationHandler());
    }
}

class MapperInvocationHandler implements InvocationHandler {
    @Override
    public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
        System.out.println("before");
        // method.getAnnotation(Select.class);
        System.out.println("after");
        return "success";
    }
}
```
- 既然如此那么猜测查询逻辑 连接数据库逻辑 参数解析逻辑都在 InvocationHandler 中应该完成
