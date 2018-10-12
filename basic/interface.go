package main

import "fmt"
//接口和接口隐式实现，接口的值包括(value, type)  value是值 type类型
type Phone interface {
	call()
}
type OnePhone struct {

}
type TwoPhone struct {

}

func (onePhone OnePhone) call()  {
 fmt.Println("onePhone实现的方法----------------------")
}
func (twophone TwoPhone) call()  {
	fmt.Println("TwoPhone实现的方法+++++++++++++++++++++++")
}
func main()  {
 var phone Phone
 phone= new(OnePhone)
 phone.call()
 phone=new(TwoPhone)
 phone.call()


 var i I;
 t:=T{"TTTTTT"}
 t.M()
 fmt.Println(t.S) //因为是值接收，所以该步没有效果

 i=T{"ddsfsdfsfsfs"}
 describe(i)
}

type I interface {
	M()
}
type T struct {
	S string
}

func (t T) M()  {
	t.S="fffffffff" //该步骤只在该方法里有效
	fmt.Println(t.S)
}

type F float64

func (f F) M()  {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}