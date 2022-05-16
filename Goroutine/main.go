package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		//只使用"i"会出现打印多次i=4的情况，因为Golang中允许启动的协程中引用外部的变量
		//使用一个临时变量
		io := i
		go func() {
			fmt.Println(io)
		}()
		//或者
		/*
			for i:=0;i<10;i++{
				go func(i0 int){
					fmt.Println(i0)
				}(i)
			}
		*/
	}
	fmt.Println(10)
	time.Sleep(time.Second)
}
