package main

import (
	"fmt"
	"sync"
	"time"
)

var(
	x int64
	wg sync.WaitGroup
	lock sync.Mutex
	rwlock sync.RWMutex
)

func write(){
	rwlock.Lock()//加写锁
	x=x+1
	time.Sleep(10*time.Millisecond)
	rwlock.Unlock()
	wg.Done()
}

func read(){
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
	wg.Done()
}

func main(){
	start:=time.Now()
	for i:=0;i<10;i++{
		wg.Add(1)
		go write()
	}
	for i:=0;i<1000;i++{
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end:=time.Now()
	fmt.Println(end.Sub(start))
}