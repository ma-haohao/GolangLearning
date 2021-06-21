package main
import "fmt"
func main(){
	//1. 取地址
	//2. 根据地址取值
	var b []int
	b=append(b,10)
	c:=&b
	fmt.Printf("%v\n",&b)
	fmt.Printf("%v\n",c)
	(*c)[0]=15
	fmt.Printf("%T\n",len(*c))
	fmt.Printf("%v\n",b)
	n:=18
	a:=&n
	fmt.Println(&n)
	fmt.Println(a)
	fmt.Printf("%T\n",a)
	m:=*a //*a表示*int指针指向的值
	*a=19
	fmt.Printf("%T\n",m)
	fmt.Println(*a)
	fmt.Println(n)
}
/*func main(){
	//var a *int //未初始化，为一个nil
	//new申请一个内存地址
	var a=new(int)
	*a=100
	fmt.Println(a)
	var a2 *int
	fmt.Println(a2)
}*/