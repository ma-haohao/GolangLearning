package main

import "fmt"
//break & continue
func main(){
	/*for i:=0;i<10;i++{
		if i==5{
			continue
		}
		fmt.Println(i)
	}*/
	var n=3
	switch n{
	case 1:fmt.Println("daMuZhi")
	case 2:fmt.Println("shiZhi")
	case 3:fmt.Println("zhongZhi")
	default:
		fmt.Println("error")
	}
}

