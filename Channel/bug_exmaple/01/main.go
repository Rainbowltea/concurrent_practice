package main

import (
	"fmt"
	"sync"
)

/*
	问题：23行并不会在接受完chan中的值后退出循环
	应改为 task,ok:=<-ch，判断ok的bool值
*/
func main() {
	wg := sync.WaitGroup{}

	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	wg.Add(3)
	for j := 0; j < 3; j++ {
		go func() {
			for {
				task := <-ch
				// 这里假设对接收的数据执行某些操作
				fmt.Println(task)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
