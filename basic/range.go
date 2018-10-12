package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
//for循环返回的俩值 一个是数组下标，一个是数组下标对应的值
func main ()  {
	for i,v :=range pow{
		fmt.Println(i,v)
	}

	fmt.Println("---------------------------------")
	pow2 :=make([]int, 10)
	for i := range pow2{
		pow2[i]= 1 << uint(i) //二进制1向左移i位 相当于 2的i次方
	}
	for _,v := range pow2{
		fmt.Println(v)
	}
}
