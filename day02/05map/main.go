package main

import "fmt"
	//"math/rand"
	//"sort"
	//"time"

/*func main(){
	scoreMap:=make(map[string]int,8)
	scoreMap["张三"]=90
	scoreMap["小明"]=100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n",scoreMap)
}*/
/*func main(){
	userInfo:=map[string]string{
		"username":"沙河小王子",
		"password":"123456",
	}
	fmt.Println(userInfo)
}*/
/*func main(){
	scoreMap:=make(map[string]int,8)
	scoreMap["张三"]=90
	scoreMap["小明"]=100
	v,ok:=scoreMap["小范"]
	if ok {
		fmt.Println(v)
	}else{
		fmt.Println(ok)
		fmt.Println("查无此人")
	}
}*/

/*func main(){
	scoreMap:=make(map[string]int)
	scoreMap["张三"]=90
	scoreMap["李四"]=100
	scoreMap["小范"]=80
	for k,v:=range scoreMap{
		fmt.Println(k,v)
	}
	for k:=range scoreMap{
		fmt.Println(k)
	}
	//使用delete()函数来删除键值对
	delete(scoreMap,"张三")
	fmt.Println(scoreMap)
}*/

/*func main(){
	rand.Seed(time.Now().UnixNano())//初始化随机种子

	var scoreMap=make(map[string]int,200)

	for i:=0;i<100;i++{
		key:=fmt.Sprintf("stu%02d",i)
		value:=rand.Intn(100)
		scoreMap[key]=value
	}
	//取出map中的所有key存入切片keys
	var keys=make([]string,0,200)
	for key:=range scoreMap{
		keys=append(keys,key)
	}
	//对切片进行排序
	sort.Strings(keys)
	for _,key:=range keys{
		fmt.Println(key,scoreMap[key])
	}
}*/

func main(){
	var sliceMap=make(map[string][]string,3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key:="中国"
	value,ok:=sliceMap[key]
	if !ok{
		value=make([]string,0,2)
	}
	value=append(value,"北京","上海")
	sliceMap[key]=value
	fmt.Println(sliceMap)
}