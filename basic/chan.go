package main

import "fmt"

//信道

func main()  {
	//s := []int{7, 2, 8, -9, 4, 0}
	//
	//c := make(chan int)
	//go sum(s[:len(s)/2], c)
	//go sum(s[len(s)/2:], c)
	//x, y := <-c, <-c // 从 c 中接收
	//
	//fmt.Println(x, y, x+y)


	s := []int{7, 2, 8, -9, 4, 0}
	c :=make(chan  int)
	go sum(s,c)
	go sum(s[len(s)/2:],c)
	x,y :=  <- c, <- c
	fmt.Println(x,y)


	ch :=make(chan  int ,3)
	ch <- 2
	ch <-4
	fmt.Println(<-ch)
	fmt.Println(<-ch)


	cc := make(chan int, 10)
	go fabonaqie(cap(cc), cc)
	for i := range cc {
		fmt.Println(i)
	}
}

func sum(s []int,c chan int)  { //int类型的信道
 sum := 0
 for _,v :=range s{
	 sum += v
 }
 c <- sum
}

func fabonaqie(n int,c chan int)  {
	x,y :=0,1
	for i:=0;i<n;i++{
		c <- x
		x,y = y,x+y
	}
	close(c)
}