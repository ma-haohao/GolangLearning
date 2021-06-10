package main
import "fmt"

func main(){
	//基本格式
	for i:=0;i<10;i++{
		fmt.Println(i)
	}
	//变种1
	i:=5
	for ;i<10;i++{
		fmt.Println(i)
	}
	//变种2
	for ;i<10;{
		fmt.Println(i)
		i++
	}
	//for range(键值循环)
	s:="hello"
	for _,v:=range s{
		fmt.Printf("%c\n",v)
	}
}