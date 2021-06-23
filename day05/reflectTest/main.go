package main

import (
	"encoding/json"
	"fmt"
)

type Cat struct{
	
}

type person struct{
	Name string `json:"name"`
	Age int `json:"age"`

}

func main(){
	str:=`{"name":"zhoulin","age":9000}`
	var p person
	json.Unmarshal([]byte(str),&p) //运行到这里才去判断类型
	fmt.Println(p.Name,p.Age)
}