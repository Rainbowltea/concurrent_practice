package main

import "fmt"

func recv(c chan int) {
	ret := <-c
	fmt.Println(ret)
}
func main() {
	//出现死锁，需要另开一个线程来接受无缓冲channel的传值
	// ch1 := make(chan int)
	// ch1 <- 10
	// x := <-ch1
	// close(ch1)
	// fmt.Println(x)
	ch1 := make(chan int)
	go recv(ch1)
	ch1 <- 10
	fmt.Println("success")
}
