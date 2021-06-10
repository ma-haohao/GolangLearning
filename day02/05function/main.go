package main
import "fmt"

func sayHello(){
	fmt.Println("hello function")
}
func intSum(x int, y int)int{
	return x + y
}
//多返回值
func calc(x,y int)(int,int){
	sum:=x+y
	sub:=x-y
	return sum,sub
}
//返回值命名
func calc1(x,y int)(sum,sub int){
	sum=x+y
	sub=x-y
	return
}
//返回值补充
func someFunc(x string)[]int {
	if x==""{
		return nil
	}
	ret:=[]int{1,2,3}
	return ret
}

func main(){
	sayHello()
	ret:=intSum(10,20)
	fmt.Println(ret)
	test1,test2:=calc(5,6)
	fmt.Println(test1,test2)
	test3,test4:=calc1(99,8)
	fmt.Println(test3,test4)
	ret1:=someFunc("test")
	fmt.Println(ret1)
}
