package main

import (
	"fmt"
)

//Person结构体
type person struct{
	name,city string
	age int8
}

type Person struct{
	string
	int
}

//Address 地址结构体
type Address struct{
	Province string
	City string
}
//User 结构体
type User struct{
	Name string
	Gender string
	Address Address
}

/*func main(){
	user1:=User{
		Name:"zhangsan",
		Gender: "man",
		Address: Address{
			Province: "山东",
			City: "威海",
		},
	}
	fmt.Printf("user-1: %#v\n",user1)
}*/

//NerPerson 构造函数
func newPerson(name,city string, age int8) *person{
	return &person{
		name:name,
		city:city,
		age:age,
	}
}

//Dream Person的方法
func(p person) Dream(){
	fmt.Printf("%s的梦想是好好学习！\n",p.name)	
}

func(p *person) SetAge(newAge int){
	p.age=int8(newAge)
}
func(p person) SetAge2(newAge int){
	p.age=int8(newAge)
}

type Animal struct{
	name string
}

func (a *Animal)move(){
	fmt.Printf("%s会动!\n",a.name)
}

type Dog struct{
	Feet int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang(){
	fmt.Printf("%s会汪汪汪~\n",d.name)
}

/*func main(){
	p9:=newPerson("zhangsan","shanghai",18)
	fmt.Printf("%#v\n",p9)
	p10:=newPerson("小王子","上海",25)
	p10.Dream()
	fmt.Println(p10.age)
	p10.SetAge(30)
	fmt.Println(p10.age)
	p10.SetAge2(40)
	fmt.Println(p10.age)
	p1:=Person{
		"小王子",
		18,
	}
	fmt.Printf("%#v\n",p1)
	fmt.Println(p1.string,p1.int)
}*/

/*func main(){
	d1:=&Dog{
		Feet: 4,
		Animal: &Animal{
			name: "LELE",
		},
	}
	d1.wang()
	d1.move()
}*/

type Student struct{
	ID int
	Gender string
	Name string
}

type Class struct{
	Title string
	Students []*Student //切片，传入的是指向student结构体的指针
}

/*func main(){
	c:=&Class{
		Title: "101",
		Students: make([]*Student,0,200),
	}
	for i:=0;i<10;i++{
		stu:=&Student{
			Name: fmt.Sprintf("stu%02d",i),
			Gender: "Male",
			ID:i,
		}
		c.Students=append(c.Students, stu)
	}
	//JOSON序列化：结构体->JSON格式字符串
	data,err:=json.Marshal(c)
	if err!=nil{
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n",data)
	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1:=&Class{}
	err=json.Unmarshal([]byte(str),c1)
	if err!=nil{
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n",c1)
}*/

type Person2 struct{
	name string
	age int8
	dreams []string
}

func (p *Person2) SetDreams(dreams []string){
	p.dreams=dreams
}

func main(){
	p1:=Person2{name:"xiaowang",age:18}
	data:=[]string{"吃饭","睡觉","打豆豆"}
	p1.SetDreams(data)

	data[1]="不睡觉"
	fmt.Println(p1.dreams)

}