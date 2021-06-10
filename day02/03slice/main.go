package main
import "fmt"
/*func main(){
	//切片的类型声明
	var a []string
	var b=[]int{}
	var c=[]bool{false,true}
	var d=[]bool{false,true}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a==nil)
	fmt.Println(b==nil)
	fmt.Println(c==nil)

}*/
/*func main(){
	a:=[5]int{1,2,3,4,5}
	s:=a[1:3] //闭开区间
	a[1]=9 //s的切片本质也是一个新的指针，而不是复制了a的范围，切片的上限确定了len的长度
	fmt.Printf("s:%v len(s): %v cap(s): %v\n",s,len(s),cap(s))
	s2:=s[3:4]
	fmt.Printf("s2:%v len(s): %v cap(s): %v\n",s2,len(s2),cap(s2))
}*/
/*func main(){
	a:=[5]int{1,2,3,4,5}
	t:=a[1:3:5]
	fmt.Printf("s:%v len(s): %v cap(s): %v\n",t,len(t),cap(t))
}*/
/*func main(){
	a:=make([]int,3,10)
	a[0]=1
	a[1]=3
	a[2]=5
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	for i:=0;i<len(a);i++{
		fmt.Println(i,a[i])
	}
	for index,value:=range a{
		fmt.Println(index,value)
	}
}*/
/*func main(){
	var s []int
	s=append(s,1)
	s=append(s,1,2,3,4)
	s2:=[]int{5,6,7}
	s=append(s,s2...)
	fmt.Println(s2)
	fmt.Println(s)
}*/
//使用copy()函数来复制切片
func main(){
	a:=[]int{1,2,3,4,5}
	c:=make([]int,5,5)
	copy(c,a)
	fmt.Println(c)
	a[1]=10000
	fmt.Println(a)
	fmt.Println(c)
}