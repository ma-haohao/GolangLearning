package main

import "fmt"


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
}