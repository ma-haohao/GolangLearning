package main

import "fmt"

//定义全局变量
var num int64=10

func testGlobalVar(){
	fmt.Printf("num%d\n",num)
	fmt.Printf("Typeof the var: %T\n",num)
}
func main(){
	testGlobalVar()
}