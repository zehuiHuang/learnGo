package main

import "fmt"

//数组表达式 var a [10] int
func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)


	//切片  左边从0开始，右边不包括，冒号分开
	primes2 := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes2[1:4]
	fmt.Println(s)
	fmt.Println(primes)

	//切片并不存储数据，只是引用底层数组的一部分,如果分片改了数据，原来的数据也会改变
	names :=[4]string{"a","b","c","d"}
	fmt.Println(names)
	ab := names[0:2]
	cd :=names[1:3]
	fmt.Println(ab,cd)
	ab[1] ="f"
	fmt.Println(ab)
	fmt.Println(names)





}
