package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type IpAddress []byte
//Person类的方法
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
//fmt.Sprintf()  打印并返回打印的字符串
func main() {
	a := Person{"Arthur Dent", 42}
	//z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a.String())
	//fmt.Println(a, z)


	host := map[string]IpAddress{
		"yiyi":{1,2,3,4,5},
		"tiny":{6,7,8,9},
	}
	for _,i :=range host{//_表示键的省略
		fmt.Println(i)
	}
}
