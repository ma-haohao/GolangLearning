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