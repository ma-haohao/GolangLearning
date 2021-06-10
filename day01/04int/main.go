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