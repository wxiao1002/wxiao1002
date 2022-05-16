- 标识符:</p>
  包括常量 变量 类型 函数 结构体等等，如果以一个大写字母开头，比如Group1,那么使用这种形式的标识符对象就可以被外部包的代码使用（外部包需要导入这个包），这辈成为导出 类似对象语言的public,</p>
  如果标识符号以 小写字母开头，则对包外是不可见的
- 注意： </br>
  { 不能单独成一行  </br>
- 关键字 </br>
break default func interface select case defer go map struct chan else goto package switch const fallthrough if range type continue for import return var 
- 预定义标识符 </br>
append bool byte cap close complex complex64 complex128 uint 16 copy false float32 float64 image int int8 int16 int32 .....
 
- 变量声明 </br>

  var age int;

- 格式化字符串 </br>
  fmt.Sprintf

- 数据类型
#### 布尔型
var a bool = true
#### 数字类型
整数类型 int 与浮点类型float32,float64 支持复数，位运算采用补码
#### 字符串类型 
go 的字符串是单个字节连接起来，使用UTF-8编码标识Unicode 文本
#### 派生类型
 - 指针 
 - 数据
 - 结构化
 - channel
 - 函数类型
 - 切片类型
 - 接口类型
 - Map 类型
##### 数字类型
- uint8 无符号8位整数 0-255
- uint16 无符号16位整数 0-65535
- uint32 无符号32位整数
- uint64 无符号 64 位整数
- int8 有符号8位整数 -128～127 <br/>
....

<p>浮点型</p>
- float32 32位浮点型数
- float64 54位浮点型
- complex64 32位实数和虚数
- complex128 64位实数和虚数

<p>其他的数字类型</p>
- byte 类似uint8
- rune 类似 int32
- uint 32或者64
- int 与uint 一样大小
- uintptr 无符号整型 用于存放一个指针

## go 语言变量
var identifier type <br/>
一次声明多个变量：var identifier1, identfier2 type
<p>数值类型 默认值是0 bool 默认是 false，字符串默认是 "" ，以下类型是nil</p>
- var a *int 
- var a []int
- var a map[string] int
- var a chan int 
- var a func(string) int
- var a error  // error 是接口

### 声明并赋值，如果已经声明过则会报错
v_name := v_value

### 多变量赋值
var name1, name2, name3 type
name1, name2, name3 = va1, va2,va3
</br>
var name1, name2 = var1, var2 // 会自动推断类型 <br/>
name1, name2 := value1, value2

### 全局变量的一般写法
var (
 name1 type
 name2 type
)
### 值类型与引用类型
所有像int,float, bool string 这些基本类型属于值类型，这种类型的变量直接指向内存在内存中的值
<p>使用符号 = 将一个变量的值赋值给另一个变量时，如 j =i ,实际上是在内存中将i值进行了拷贝</p>
可以使用&i 来获取变量i的内存地址，
<p>**值类型变量的值存储在堆上</p>
引用类型变量r1 存储的是r1 所在的内存地址，或者内存地址中第一个字所在位置</p>
这个内存地址被称为**指针**，这个指针实际上也被存储在另外一个指中，比如r2 = r1,只有引用地址被复制

### 空白标识符 _ 
被用于抛弃值，如 5 在 _,b = 5, 7 中被抛弃，_ 实际上是一个只写变量，你不能得到他的值，这样做的原因是 你必须使用所有声明的变量，又是你并不需要使用从一个函数得到所有返回值
