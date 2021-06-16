package main

import "fmt"

func show(a interface{}){
	fmt.Printf("type %T value: %v\n",a,a)
}

func main(){
	//定义一个空接口
	var a interface{}
	s:="hello shahe"
	a=s
	fmt.Printf("type: %T value: %v\n",a,a)
	i:=100
	a=i
	fmt.Printf("type: %T value: %v\n",a,a)
	j:=false
	a=j
	fmt.Printf("type: %T value: %v\n",a,a)
	var studentInfo=make(map[string]interface{})
	studentInfo["name"]="xiaowang"
	studentInfo["age"]=18
	studentInfo["married"]=false
	studentInfo["friends"]=[]string{"xiaofan","xiaoma","xiaowu"}
	fmt.Println(studentInfo)
	show(false)
	show("test")
	show(10)

	var x interface{}
	x="hello shahe"
	v,ok:=x.(string)
	if ok{
		fmt.Println(v)
	}else{
		fmt.Println("类型断言失败")
	}
}