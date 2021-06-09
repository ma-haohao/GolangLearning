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