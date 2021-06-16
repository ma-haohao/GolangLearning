package main

import "fmt"
//类型定义
type NewInt int
//类型的别名
type MyInt=int

//person结构体
type person struct{
	name,city string
	age int8
}

/*func main(){
	var a NewInt
	var b MyInt
	fmt.Printf("Type of a: %T\n",a)
	fmt.Printf("Type of a: %T\n",b)
	var p1 person
	p1.name="shahe"
	p1.city="beijing"
	p1.age=18
	fmt.Printf("p1=%#v\n",p1)
	var user struct{Name string;Age int}
	user.Name="xiao wang zi"
	user.Age=18
	fmt.Printf("%#v\n",user)
	var p2=new(person)
	p2.name="test"
	fmt.Printf("p2=%#v\n",p2)
}*/
//取结构体的地址实体化
type test struct{
	a int8
	b int8
	c int8
	d int8
}
func main(){
	p3:=&person{}
	fmt.Printf("%T\n",p3)
	fmt.Printf("p3=%#v\n",p3)
	p3.name="test"
	p3.age=18
	p3.city="shanghai"
	fmt.Printf("p3=%#v\n",p3)
	n:=test{
		1,2,3,4,
	}
	fmt.Printf("n.a:%p\n",&n.a)
	fmt.Printf("n.b:%p\n",&n.b)
	fmt.Printf("n.c:%p\n",&n.c)
	fmt.Printf("n.d:%p\n",&n.d)

}