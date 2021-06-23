package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func producer(ch1 chan<- string,pNum int){
	for i:=0;i<10;i++{
		content:=fmt.Sprintf("item %v from p%v", i,pNum)
		fmt.Println(content)
		ch1<-content
	}
	wg.Done()
}

func customer(ch1 <-chan string,cNum int){
	for i:=0;i<10;i++{
		v:=<-ch1
		fmt.Printf("c%v received: %v\n",cNum,v)
	}
	wg.Done()
}

func main(){
	ch:=make(chan string,5)
	for i:=1;i<=5;i++{
		wg.Add(1)
		go producer(ch,i)
		wg.Add(1)
		go customer(ch,i)
	}
	wg.Wait()
}