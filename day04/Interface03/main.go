package main

import "fmt"

type animal interface{
	move()
	eat(sth string)
}

type cat struct{
	name string
	feet int8
}
func (c cat)move(){
	fmt.Println("猫动！")
}
func (c cat)eat(sth string){
	fmt.Printf("%s吃%s\n",c.name,sth)
}

type chicken struct{
	name string
	feet int8
}

func (c chicken)move(){
	fmt.Println("鸡动！")
}
func (c chicken)eat(sth string){
	fmt.Printf("%s吃%s\n",c.name,sth)
}


func main(){
	var a1 animal //定义一个接口类型的变量
	bc:=cat{
		"淘气",
		4,
	} //定义一个cat类型的变量
	a1=bc
	fmt.Println(a1)
	a1.eat("老鼠")
	fmt.Printf("%T\n",a1)
}