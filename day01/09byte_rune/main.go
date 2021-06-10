package main

import (
	"fmt"
)
func main(){
	s2:="白萝卜"
	s3:=[]rune(s2)//把字符串强制转换成rune切片
	s3[0]='红'
	fmt.Println(string(s3))
	c1:="红"
	c2:='红'
	fmt.Printf("%T\n",c1)
	fmt.Printf("%T\n",c2)

}

