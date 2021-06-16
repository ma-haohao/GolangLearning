package main

import (
	"errors"
	"fmt"
	"strings"
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

func adder() func(int) int{
	var x int
	return func(y int) int{
		x+=y
		return x
	}
}

func adder2(x int) func(int) int{
	return func(y int) int {
		x+=y
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string{
	return func(name string)string {
		if !strings.HasSuffix(name,suffix){
			return name+suffix
		}
		return name
	}
}

func calc1(base int)(func(int) int,func(int) int){
	add:=func(i int) int{
		base+=i
		return base
	}
	sub:=func(i int) int{
		base-=i
		return base
	}
	return add,sub
}

func main(){
	/*var c calculation
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
	}(100,99)*/

	f:=adder()
	fmt.Println(f(10))
	fmt.Println(f(20))
	fmt.Println(f(30))
	f1:=adder()
	fmt.Println(f1(40))
	fmt.Println(f1(50))

	var f2=adder2(10)
	fmt.Println(f2(10))
	fmt.Println(f2(50))

	jpgFunc:=makeSuffixFunc(".jpg")
	txtFunc:=makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test"))
	fmt.Println(txtFunc("test"))
	
	f10,f20:=calc1(10)
	fmt.Println(f10(1),f20(2))
	fmt.Println(f10(9),f20(3))


}