package main
import "fmt"
//常量
//定义了常量之后不能修改
const pi=3.14159
//批量申明常量
const (
	STATUSOK=200
	NOTFOUND=404
)
//批量申明常量时，如果某一行没有赋值，默认和上一行一致
const (
	n1=100
	n2
	n3
)

//iota
const(
	a1=iota
	a2=100
	_
	a4=iota
)

func main(){
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println("a1:",a2)
	fmt.Println("a4:",a4)
}