package main

import "fmt"

//	类型断言，相当于java的instanceof判断类型
//一般用于switch语句中

//value:=em.(T)  其中em必须时interface类型才可以
func main()  {
	var i interface{}="jerry"

	s := i.(string)
	fmt.Println(s)
	s,ok :=i.(string)
	fmt.Println(s,ok)
	do("ffff")
	do(123123)
}

func do( i interface{})  {
	switch v :=i.(type) {

	case int:
		fmt.Println("int类型，值是：",v)
	case string:
		fmt.Println("字符串类型："+v)
	default:
		fmt.Println("未知类型！！")

	}

}