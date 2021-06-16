package main
import "fmt"

//定义一个能叫的类型
type Sayer interface{
	Say() //只要实现了say方法的变量，都是speak类型,方法签名
}

type Cat struct{}

func (c Cat) Say(){
	fmt.Println("miaomiaomiao")
}

type Dog struct{}

func (d Dog) Say(){
	fmt.Println("wangwangwang")
}

func da(x Sayer){
	x.Say() //打了就叫
}

func main(){
	var c Cat
	var d Dog
	da(c)
	da(d)
}