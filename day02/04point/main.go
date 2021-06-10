package main
import "fmt"
/*func main(){
	//1. 取地址
	//2. 根据地址取值
	n:=18
	a:=&n
	fmt.Println(&n)
	fmt.Println(a)
	fmt.Printf("%T\n",a)
	m:=*a //*a表示int类型的指针
	fmt.Printf("%T\n",m)
}*/
func main(){
	//var a *int //未初始化，为一个nil
	//new申请一个内存地址
	var a=new(int)
	*a=100
	fmt.Println(a)
	var a2 *int
	fmt.Println(a2)
}