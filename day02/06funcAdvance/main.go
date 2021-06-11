package main

import (
	"errors"
	"fmt"
)

//定义全局变量
var num int64=10

func testGlobalVar(){
	fmt.Printf("num%d\n",num)
	fmt.Printf("Typeof the var: %T\n",num)
}

func testLocalVar2(x,y int){
	fmt.Println(x,y)
	if x>0{
		z:=100
		fmt.Println(z) //变量z只有在if语句块内才能生效
	}
	//fmt.Println(z)
}

type calculation func(int,int) int
func add(x,y int) int{
	return x+y
}
func sub(x,y int) int{
	return x-y
}


func add1(x,y int)int{
	return x+y
}
//函数作为参数
func calc(x,y int,op func(int,int)int)int{
	return op(x,y)
}
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

func testLocalVar3(){
	for i:=0;i<10;i++{
		fmt.Println(i)
	}
	//fmt.Println(i) //i只能在for语句块内使用
}

func main(){
	var c calculation
	c=add
	fmt.Printf("Type of c:%T\n",c)
	fmt.Println(c(10,20))
	ret2:=calc(10,90,add1)
	fmt.Println(ret2)
	calMethod,_:=do("+")
	ret3:=calMethod(90,100)
	fmt.Println(ret3)
	add:=func(x,y int){
		fmt.Println(x+y)
	}
	add(10,110)
	func(x,y int){
		fmt.Println(x-y)
	}(100,99)
}