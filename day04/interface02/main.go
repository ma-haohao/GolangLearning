package main
import "fmt"

//接口示例2
//不管是什么牌子的车，都能跑

type car interface{
	run()
}

type falali struct{
	brand string
}

func (f falali) run(){
	fmt.Printf("%s速度70迈\n",f.brand)
}

type baoshijie struct{
	brand string
}
func (b baoshijie) run(){
	fmt.Printf("%s速度700迈\n",b.brand)
}

func drive(c car){
	c.run()
}

func main(){
	var f1=falali{
		"法拉利",
	}
	var b1=baoshijie{
		"宝时捷",
	}
	drive(f1)
	drive(b1)
	var ss car //定义一个接口类型的变量
	ss=f1
	ss=b1
	drive(ss)
}