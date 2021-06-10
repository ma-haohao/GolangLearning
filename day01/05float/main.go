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
	//f2的值不能赋值给f1，类型不同
} 
 