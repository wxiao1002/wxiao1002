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

### 常量
常量是一个简单值的标识符，在程序运行时，不会被修改的量<br/>
常量中的数据类型只可以是 布尔、数字（整形、浮点、复数）和字符 <br/>
const identifier type = value
```go

const b string = "b"
const c = "c"
```
常量也可以作为枚举类
```go
const (
	Unknow =0
	Female =1
	Male =2
)
```
iota，特殊常量，可以认为是一个可以被编译器修改的常量
```go
const (
	a = iota
	b = iota
)
```
当第一个iota 等于0 每当iota在新的一行被使用时候，他的值都会自动加1
```go
// a=0, b=1, c=2 简写成如下格式
const (
	a = iota 
	b 
	c
)
// iota 表示从 0 开始自动加 1，所以 i=1<<0, j=3<<1（<< 表示左移的意思），即：i=1, j=6，这没问题，
// 关键在 k 和 l，从输出结果看 k=3<<2，l=3<<3。
```

### 运算符
- 算术运算符
- 关系运算符
- 逻辑运算符
- 位运算符
- 复制运算符
- 其他运算符

#### 算数运算符
```
+ - * / % ++ --
```

#### 关系运算符
```
== 检查相等 !=  > < >= <=
```

#### 逻辑运算符
```
&& 逻辑and
|| 逻辑OR
! NOT
```
#### 位运算
```
&: 0&1 = 0 ,两边为1 结果才是1
|: 1&0 = 1， 两边为0 结果才为0
^：异或 相异为1
<< 左移，左移n 相当于乘以2的n次方
>> 右移 除以2的n次方
```
#### 赋值运算符
```
+= -= >>= <<= %= |= ....
// 所有的赋值运算符 都等于 运算后再赋值
```
#### 其他运算符
```
& 返回存储地址，&a 给出变量a 的实际地址
* 指针变量 *a 是一个指针变量
```

### 优先级
```
// 依次往下 优先级最低
* / % << >> & ^
+ - | ^
== != >= 
&& || 
```

## 条件语句
```
if 

if else
switch
select // 类似switch ，但会随机选择一个可运行的case,如果没有case 则会阻塞
```
**没有三元运算符**

## 循环语句
- for 循环
```go
var i int
for i=0;i<5;i++{
	//
}
```
- 循环嵌套
### 循环控制语句
- break 跳出当前for 循环或者switch
- continue 跳过当前循环的剩下语句，执行下一轮循环
- goto 将控制转移到被标记的语句
### 无限循环
```go
if true{
	fmt.Println("aaaa")
}
```

## 函数
```
func function_name( [parameter list] ) [return_types] {
   函数体
}
```
- func: 函数开始声明
- func_name: 函数名称
- parameter list： 参数列表，参数就像一个占位符，当函数被调用时，你就可以将值传递给函数，这个值也被成为实际参数。参数列表包括参数类型、顺序、个数
- return_types：返回类型，返回一列值，无需返回不必填写

### 函数参数
函数如果使用参数，该变量成为函数的形参，形参数相当于定义在函数体内的变量<br/>
- 值传递：在调用函数时将实际参数的值赋值一部分到函数中，函数中的更改并不会影响实际参数
- 引用传递：在函数调用时将实际参数的地址传递在函数中，函数中对参数的更改将会影响到实际参数
<p>默认情况下Go语言使用的是值传递</p>

### 函数的用法
- 函数作为另一个函数的实参：函数定义后作为另一个函数的实际参数传入
```go
// 声明函数变量
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	// 使用函数
	fmt.Println(getSquareRoot(9))
```
- 闭包：匿名函数，在动态编程中使用
```go

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
    nextNumber := getSequence()
    fmt.Println(nextNumber())
}
```
- 方法：一个包含了接受者的函数,接受者是命名类型或者结构体类型的一个值或者指针
```go
package main

import "fmt"

// 定义结构体
type Circle struct {
	radius float64
}

// 该方法数据 Circle 结构体中的方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
func main() {
	var c1 Circle
	c1.radius = 20.00
	fmt.Println(c1.getArea())
}
```

## 变量作用域
作用域为已经声明的标识符所表示的常量、变量、类型和函数或包在源码中的作用范围<br/>
- 函数体内定于的变量为局部变量
- 函数外定义的变量为全局变量
- 函数定义中的变量为形式参数
**局部变量与全局变量名称可以相同，但是函数体中局部变量会被优先考虑**

## 数组
Go 语言提供数组类型的数据结构，数组是具有相同类型的一族长度固定的数据序列
```
var variable_name [SIZE] variable_type
```
定义一个一维数组 长度为10 数据类型是 float
```go
var blance [10] float32
```
### 数组的初始化
```go
var balance  = [5] float32 {100.0,2.0,3.4,7.0,50.0}
```
数组长度不确定，使用... 代替长度，编译器会根据元素个数自行推断
```go
cc := [...]int{1, 2, 3}
	fmt.Println(cc[0])
```
### 多维数组
声明
```go
var var_name [SIZE][SIZE]...[SIZE] var_type
```
声明一个三维整型数组
```go
var threed [3] [3] [3] int 
```
多维数组中添加一维数组 **append**
```go
	// 创建数组
	values := [][]int{}
	// 创建1维数组
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	// 向二维数组中添加一维数组
	values = append(values, row1)
	values = append(values, row2)
	fmt.Println(values[1][1])
```
声明并初始化数组
```go
a := [3,4]  init {
 {0,1,2,3},
 {4,5,6,7},
 {8,9,10,11}
}
```
循环输出 **range**
```go
for i := range sites {
		fmt.Println(sites[i])
}
```
### 向函数中传递数组
在函数中传递数组，只需要将形参设置成数组
```go
func getAverage ( arr [] int, size int) float32{
	var i int
	var sum float32
	for i=0;i<size;i++ {
		sum += arr[i]
    }
	avg := sum / size
	return avg
	
}
```

## 指针
变量是一种是哟过方便的占位符，用于引用计算机内存地址，**取地址符号&，放到一个变量前使用会返回相应变量的内存地址**
```go
func main() {
	var a int = 10
	fmt.Printf("%x", &a)
}
```
一个指针变量指向了一个值的内存地址，类似于变量和常量，指针的声明如下
```
var var_name *var_type
```
var_type 为指针类型<br/>
var_name 为指针变量名，*号用于指定变量为一个指针
```go
var ip *int // 指向int 类型
var fp *float64 // 指向浮点类型
```
### 如何使用指针
- 定义指针变量
- 为指针变量赋值
- 访问指针变量中指向地址的值
```go
var c int = 20 // 声明实际变量
	var ip *int    // 声明指针
	ip = &c        // 指针变量的存储地址
	fmt.Printf("a 变量地址是%x\n", &c)

	fmt.Printf("*ip 变量的值%d", *ip)
	
```
### 空指针
当一个指针被定义后没有分配任何变量是，他的值是nil,nil指针也被称为空指针，一个指针变量通常缩写称ptr
```go
var ptr *int
fmt.Printf("%x",ptr)
```
空指针判断
```go
if(ptr != nil) // 不是空指针
```
### 指针数组
定义一个指针数组来存储地址
```
// 声明 整型 指针 数组
var ptr [] *int
```
ptr 为整形的指针数组，因此每一个元素都指向一个值
```go
    arr := [3]int{100, 200, 300}
	var i int
	var ptr [3]*int
	for i = 0; i < 3; i++ {
		ptr[i] = &arr[i]
	}
	for i = 0; i < 3; i++ {
		fmt.Printf("a[%d] is%d", i, *ptr[i])
	}

```
### 指向指针的指针
如果一个指针变量存放的是另一个指针变量的地址，则称为这个指针便可为指向指针的指针变量<br />
当定义一个指向指针的指针变量是，第一个指针存放的是第二个指针的地址，第二个指针存放的是变量的地址
```go
// 声明
var ptr **int
```
访问指向指针的指针变量时候需要两个*
```go
func main() {
 var a int
 var ptr *int
 var pptr **int
 a = 300
 ptr = &a
 pptr = &ptr
 fmt.Println(**pptr)
}
```
### 向函数传递指针参数
允许向函数中传递指针，只需要在函数定义中参数设置为指针类型即可
```go

```
