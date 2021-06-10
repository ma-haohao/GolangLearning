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