## Go语言文件的基本结构

```go
package main

//导入语句， 导入的包要用双引号来包裹
import "fmt"

//函数外只能放置标识符（变量，常量，函数，类型）的申明

//程序的入口
func main(){
	fmt.Println("Hello world!")
}
```

## 变量申明

```go
package main

import "fmt"
//申明变量
var num int

//推荐使用驼峰式的命名
//批量申明变量
var (
	name string //""
	age int  //0
	isOk bool //false
)

func foo()(int,string){
	return 10,"test"
}

func main(){
	name="lixiang"
	age=18
	isOk=true
	//var heiheihei string 非全局变量申明后必须使用，不然编译时会报错
	fmt.Printf("name:%s \n",name)
	fmt.Println(age)//打印完指定内容自动加换行符

	//申明变量同时赋值
	var s1 string="xiaowang"
	fmt.Println(s1)
	//类型推到(根据值判断对象是什么类型)
	var s2="xiaoli"
	fmt.Println(s2)
	//短变量申明,只能在函数中使用
	s3:="xiaoma"
	fmt.Println(s3)
	//匿名变量,用_来接收返回值
	_,y:=foo()
	println(y)
	//同个作用域中不能重复申明同一个变量
}
```

## Go常量

### 常量定义

```go
package main
import "fmt"
//常量
//定义了常量之后不能修改
const pi=3.14159
//批量申明常量
const (
	STATUSOK=200
	NOTFOUND=404
)
//批量申明常量时，如果某一行没有赋值，默认和上一行一致
const (
	n1=100
	n2
	n3
)
func main(){
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
}
```

### iota

是go语言中的常量计数器，只能在常量的表达式中使用

在关键词出现时将被重置为0，每新增一行iota计数一次

## Go 不同类型

### Go整形

Go本身不能定义二进制数，通过整形去转换

```go
package main

import "fmt"

//整形
func main(){
	var i1=101
	fmt.Printf("%d\n",i1)
	fmt.Printf("%b\n",i1)
	i2:=077
	fmt.Printf("%d\n",i2)
	i3:= 0x1234567
	fmt.Printf("%d\n",i3)
	fmt.Printf("%T\n",i3)
	//申明int8类型的变量
	i4:=int8(9)
	fmt.Printf("%T\n",i4)
	i5:=int16(9)
	fmt.Printf("%T\n",i5)
}
```

### 浮点型

float32与float64

```go
package main

import (
	"fmt"
)

//浮点数
func main(){
	f1:=1.2345
	fmt.Printf("%T\n",f1) //默认Go语言中的小数都是float64类型
	f2:=float32(1.233423)
	fmt.Printf("%T\n",f2) //默认Go语言中的小数都是float64类型
	//f2的值不能
```

### 布尔值

默认为false

```go
  package main

  import "fmt"
  func main(){
	  b1:=true
	  var b2 bool
	  fmt.Printf("%T\n",b1)
	  fmt.Printf("%T\n",b2)
	  fmt.Printf("%v\n",b1)
	  fmt.Printf("%v\n",b2)
  }
```

### 字符串

Go语言中的字符串必须用“ ”来包裹的

Go语言中的单引号包裹的是字符

```go
package main

import (
	"fmt"
	"strings"
)

func main(){
	s:="Hello"
	c1:='杀'
	c2:="h"
	fmt.Printf("%T\n",s)
	fmt.Printf("%T\n",c1)
	fmt.Printf("%T\n",c2)
	//多行的字符串
	//反引号内部原样输出
	s2:=`test1
test2
test3
test4
	`
	fmt.Println(s2)
	//字符串的常用操作
	fmt.Println(len(s2))
	name:="lixiang"
	age:="14"
	fmt.Println(name+age)
	ss1:=fmt.Sprintf("%s%s",name,age)
	fmt.Printf("%s\n",ss1)
	s3:="E:\\Desktop\\Test\\MaHaoYe\\HelloWorld"
	//分割
	ret:=strings.Split(s3,"\\")
	fmt.Println(ret)
	//包含
	fmt.Println(strings.Contains(ss1,"lixing"))
	fmt.Println(strings.Contains(ss1,"lixiang"))
	//前缀
	fmt.Println(strings.HasPrefix(ss1,"li"))
	//后缀
	fmt.Println(strings.HasSuffix(ss1,"xiang"))
	//子串出现的位置
	s4:="abcdeb"
	fmt.Println(strings.Index(s4,"c"))
	fmt.Println(strings.LastIndex(s4,"b"))
}
```

## 流程控制

### if 语句

```go
package main
import "fmt"

func main(){
	age:=19
	if age>35{
		fmt.Println("澳门首家线上赌场开业了！")
	}else if age>10{
		fmt.Println("青年")
	}else{
		fmt.Println("该写暑假作业了！")
	}
}
```

### for 语句

```go
package main
import "fmt"

func main(){
	//基本格式
	for i:=0;i<10;i++{
		fmt.Println(i)
	}
	//变种1
	i:=5
	for ;i<10;i++{
		fmt.Println(i)
	}
	//变种2
	for ;i<10;{
		fmt.Println(i)
		i++
	}
	//for range(键值循环)
	s:="hello"
	for _,v:=range s{
		fmt.Printf("%c\n",v)
	}
}
```

switch语句

```go
package main

import "fmt"
//break & continue
func main(){
	/*for i:=0;i<10;i++{
		if i==5{
			continue
		}
		fmt.Println(i)
	}*/
	var n=3
	switch n{
	case 1:fmt.Println("daMuZhi")
	case 2:fmt.Println("shiZhi")
	case 3:fmt.Println("zhongZhi")
	default:
		fmt.Println("error")
	}
}
```

## Go语言中的数据类型

### 数组(Array)

```go
package main

import "fmt"

func modifyArray(x [3]int){
	x[0]=100
}
func modifyArray2(x [3][2]int){
	x[2][0]=100
}

func main(){
	//数组初始化
	var testArray [3]int
	fmt.Println(testArray)
	var numArray=[3]int{1,2}
	fmt.Println(numArray)
	var cityArray=[...]string{"beijing","shanghai","chongqing"}
	fmt.Println((cityArray))
	a:=[...]int{1:1,3:5}
	fmt.Printf("%T\n",a)
	//数组遍历
	//for 循环遍历
	for i:=0;i<len(a);i++{
		fmt.Println(a[i])
	}
	//for range遍历
	for index, value:=range a{
		fmt.Println(index,value)
	}
	//多维数组
	a1:=[3][2]string{
		{"beijing","shanghai"},
		{"guangzhou","shenzheng"},
		{"chengdu","chongqing"},
	}
	fmt.Println(a1)
	//多维数组的遍历
	for _,v1:=range a1 {
		for _,v2:=range v1{
			fmt.Printf("%s\t",v2)
		}
		fmt.Println()
	}
	a3:=[3]int{10,20,30}
	modifyArray(a3)
	fmt.Println(a3)
	b:=[3][2]int{
		{1,1},
		{1,1},
		{1,1},
	}
	modifyArray2(b)
	fmt.Println(b)
}
```

### 切片(Slice)

切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。

切片是一个引用类型，它的内部结构包含`地址`、`长度`和`容量`。切片一般用于快速地操作一块数据集合。

基本声明语法

var name []T 

#### 切面的本质

切片的本质就是对底层数组的封装，它包含了三个信息：底层数组的指针、切片的长度（len）和切片的容量（cap）。

举个例子，现在有一个数组`a := [8]int{0, 1, 2, 3, 4, 5, 6, 7}`，切片`s1 := a[:5]`，相应示意图如下。<img src="https://www.liwenzhou.com/images/Go/slice/slice_01.png" alt="slice_01" style="zoom:50%;" />

切片`s2 := a[3:6]`，相应示意图如下：<img src="https://www.liwenzhou.com/images/Go/slice/slice_02.png" alt="slice_02" style="zoom:50%;" />

对数组的进一步封装，其包含了一个指向底层数组的指针，长度len() 和容量cap()

#### 切片表达式

```go
//简单的切片表达式
func main(){
	a:=[5]int{1,2,3,4,5}
	s:=a[1:3] //闭开区间
	a[1]=9 //s的切片本质也是一个新的指针，而不是复制了a的范围，切片的上限确定了len的长度
	fmt.Printf("s:%v len(s): %v cap(s): %v\n",s,len(s),cap(s))
	s2:=s[3:4]
	fmt.Printf("s2:%v len(s): %v cap(s): %v\n",s2,len(s2),cap(s2))
}
```

```go
//完整的切片表达式
func main(){
	a:=[5]int{1,2,3,4,5}
	t:=a[1:3:5]
	fmt.Printf("s:%v len(s): %v cap(s): %v\n",t,len(t),cap(t))
}
```

#### 使用make()函数构造切片

make([]T,size,cap)

```go
func main(){
	a:=make([]int,2,10)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
```

#### 判断切片是否为空

利用len(s)==0，而不是用s==nil来判断

切片共享底层数组，所以一个切片的修改会影响另一个切片的内容

#### 切片的遍历

```go
func main(){
  a:=[]int{1,3,5}
	for i:=0;i<len(a);i++{
		fmt.Println(i,a[i])
	}
	for index,value:=range a{
		fmt.Println(index,value)
	}
}
```

#### 利用append()为切片添加元素

```go
func main(){
	var s []int
	s=append(s,1)
	s=append(s,1,2,3,4)
	s2:=[]int{5,6,7}
	s=append(s,s2...)
	fmt.Println(s2)
	fmt.Println(s)
}
```

#### 使用copy()函数来复制切片

```go
	func main(){
	a:=[]int{1,2,3,4,5}
	c:=make([]int,5,5)
	copy(c,a)
	fmt.Println(c)
	a[1]=10000
	fmt.Println(a)
	fmt.Println(c)
}
```

### 指针

Go语言中不存在指针操作

1. &: 取地址

2. *: 根据地址取值

```go
package main
import "fmt"
func main(){
	//1. 取地址
	//2. 根据地址取值
	n:=18
	a:=&n
	fmt.Println(&n)
	fmt.Println(a)
	fmt.Printf("%T\n",a)
	m:=*a //*a表示int类型的指针
	fmt.Printf("%T\n",m)
}
```

我们来看一下`b := &a`的图示：

<img src="https://www.liwenzhou.com/images/Go/pointer/ptr.png" alt="取变量地址图示" style="zoom:50%;" />

### new 和 make

new是给基本数据类型分配内存的：int/bool/float等（比较少用）返回的是对应类型的指针

make是用来给slice，map，channel等申请内存的，make函数返回的是对应的这三个类型本身 

```go
package main
import "fmt"
/*func main(){
	//1. 取地址
	//2. 根据地址取值
	n:=18
	a:=&n
	fmt.Println(&n)
	fmt.Println(a)
	fmt.Printf("%T\n",a)
	m:=*a //*a表示int类型的指针
	fmt.Printf("%T\n",m)
}*/
func main(){
	//var a *int //未初始化，为一个nil
	//new申请一个内存地址
	var a=new(int)
	*a=100
	fmt.Println(a)
	var a2 *int
	fmt.Println(a2)
}
```

### map

Go语言中提供的映射关系容器为`map`，其内部使用`散列表（hash）`实现。

map是一种无序的基于`key-value`的数据结构，Go语言中的map是引用类型，必须初始化才能使用。

#### map的基本用法

```Go
//map的基本用法
package main
import "fmt"
func main(){
	scoreMap:=make(map[string]int,8)
	scoreMap["张三"]=90
	scoreMap["小明"]=100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n",scoreMap)
}
```

```Go
//声明时填充元素
func main(){
	userInfo:=map[string]string{
		"username":"沙河小王子",
		"password":"123456",
	}
	fmt.Println(userInfo)
}
```

#### 判断某一个键是否存在

```Go
func main(){
	scoreMap:=make(map[string]int,8)
	scoreMap["张三"]=90
	scoreMap["小明"]=100
	v,ok:=scoreMap["小范"]
	if ok {
		fmt.Println(v)
	}else{
		fmt.Println(ok)
		fmt.Println("查无此人")
	}
}
```

#### map的遍历

可以用for range 来进行遍历

```go
func main(){
	scoreMap:=make(map[string]int)
	scoreMap["张三"]=90
	scoreMap["李四"]=100
	scoreMap["小范"]=80
	for k,v:=range scoreMap{
		fmt.Println(k,v)
	}
	for k:=range scoreMap{
		fmt.Println(k)
	}
	//使用delete()函数来删除键值对
	delete(scoreMap,"张三")
	fmt.Println(scoreMap)
}
```

按照指定顺序来遍历map

```go
func main(){
	rand.Seed(time.Now().UnixNano())//初始化随机种子

	var scoreMap=make(map[string]int,200)

	for i:=0;i<100;i++{
		key:=fmt.Sprintf("stu%02d",i)
		value:=rand.Intn(100)
		scoreMap[key]=value
	}
	//取出map中的所有key存入切片keys
	var keys=make([]string,0,200)
	for key:=range scoreMap{
		keys=append(keys,key)
	}
	//对切片进行排序
	sort.Strings(keys)
	for _,key:=range keys{
		fmt.Println(key,scoreMap[key])
	}
}
```

#### 值为切片类型的map

```go
func main(){
	var sliceMap=make(map[string][]string,3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key:="中国"
	value,ok:=sliceMap[key]
	if !ok{
		value=make([]string,0,2)
	}
	value=append(value,"北京","上海")
	sliceMap[key]=value
	fmt.Println(sliceMap)
}
```

### 函数

#### 函数的定义

函数的定义与调用

```go
func sayHello(){
	fmt.Println("hello function")
}

func intSum(x int, y int)int{
	return x + y
}

func main(){
	sayHello()
	ret:=intSum(10,20)
	fmt.Println(ret)
}
```

#### 返回值

Go语言中通过`return`关键字向外输出返回值。

多返回值

```go
//多返回值
func calc(x,y int)(int,int){
	sum:=x+y
	sub:=x-y
	return sum,sub
}
//返回值命名
func calc1(x,y int)(sum,sub int){
	sum=x+y
	sub=x-y
	return
}
//返回值补充
func someFunc(x string)[]int {
	if x==""{
		return nil
	}
	ret:=[]int{1,2,3}
	return ret
}

func main(){
	sayHello()
	ret:=intSum(10,20)
	fmt.Println(ret)
	test1,test2:=calc(5,6)
	fmt.Println(test1,test2)
	test3,test4:=calc1(99,8)
	fmt.Println(test3,test4)
	ret1:=someFunc("test")
	fmt.Println(ret1)
}
```

### 函数进阶

#### 变量的作用域

全局变量

```go
//定义全局变量
var num int64=10

func testGlobalVar(){
	fmt.Printf("num%d\n",num)
	fmt.Printf("Typeof the var: %T\n",num)
}
func main(){
	testGlobalVar()
}
```

局部变量

1. 函数内部定义的变量无法在该函数的外部使用
2. 全局变量与局部变量重名，优先访问局部变量
3. 在if, for, switch语句上使用定义变量的方式在语句块外部无法生效

```go
func testLocalVar2(x,y int){
	fmt.Println(x,y)
	if x>0{
		z:=100
		fmt.Println(z) //变量z只有在if语句块内才能生效
	}
	//fmt.Println(z) //此处z变量无法使用
}

func testLocalVar3(){
	for i:=0;i<10;i++{
		fmt.Println(i)
	}
	//fmt.Println(i) //i只能在for语句块内使用
}
```

#### 函数的类型与变量

定义函数的类型

```go
type calculation func(int,int) int
func add(x,y int) int{
	return x+y
}

func sub(x,y int) int{
	return x-y
}

func main(){
	var c calculation
	c=add
	fmt.Printf("Type of c:%T\n",c)
	fmt.Println(c(10,20))
}
```

#### 高阶函数

函数作为参数

```go
func add1(x,y int)int{
	return x+y
}
//函数作为参数
func calc(x,y int,op func(int,int)int)int{
	return op(x,y)
}
func main(){
  ret:=calc(10,20,add1)
  fmt.Println(ret)
}
```

函数作为返回值

```go
//函数作为返回值
func do(s string)(func(int,int)int,error){
	switch s {
	case "+":
		return add,nil
	case "-":
		return sub,nil
	default:
		err:=errors.New("无法识别的操作符")
		return nil,err
	}
}
func main(){
	ret3:=calMethod(90,100)
	fmt.Println(ret3)
}
```

#### 匿名函数

函数当然还可以作为返回值，但是在Go语言中函数内部不能再像之前那样定义函数了，只能定义匿名函数。匿名函数就是没有函数名的函数，匿名函数的定义格式如下：

```go
func(参数)(返回值){
    函数体
}
```

匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数:

```go
func main() {
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}
```

匿名函数多用于实现回调函数与闭包

