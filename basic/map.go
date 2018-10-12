package main

import (
	"fmt"
)

type Vt struct {
	a , b string
}
var m map[string]Vt
func main()  {
	m=make(map[string]Vt)
	m["ddd"]=Vt{"d","ff"}

	var m2 = map[string]Vt{
		"Bell Labs": {"sdf", "sdf"},
		"Google":    {"d", "fff"},
	}

	fmt.Println(m["ddd"])
	fmt.Println(m2)
	//判断一个map是否有一个键存在用法 v表示值，ok表示true or false
	v, ok := m2["Google"]
	//赋值和查找和java的map一样  删除的时候
	delete(m, "ddd")
	fmt.Println(v,ok)
}
