package main

import (
	"container/heap"
	"fmt"
)

type myHeap [][]int

func (h myHeap) Len() int { return len(h) }  // 绑定len方法,返回长度

func (h myHeap) Less(i, j int) bool {  // 绑定less方法
	return h[i][1] < h[j][1]  // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}

func (h myHeap) Swap(i, j int) {  // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *myHeap) Pop() interface{} {  // 绑定pop方法，从最后拿出一个元素并返回
	x:=(*h)[len(*h)-1]
	*h=(*h)[0:len(*h)-1]
	return x
}

func (h *myHeap) Push(x interface{}) {  // 绑定push方法，插入新元素
	*h = append(*h, x.([]int))
}

func main(){
	h:=&myHeap{}
	heap.Init(h)
	heap.Push(h,[]int{1,2})
	heap.Push(h,[]int{99,0})
	heap.Push(h,[]int{8,9})
	heap.Push(h,[]int{36,3})
	heap.Push(h,[]int{7,15})
	for h.Len()>0{
		val:=heap.Pop(h)
		fmt.Println(val)
	}
}