package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
//方法也可以当作参数
func  main()  {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))


	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())//调用MyFloat的方法


	//v := Vertex{3, 4}
	//v.Score(10)
	//fmt.Println(v.Y)


	//5、
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.X)
}



func compute(fn2 func(float64, float64) float64) float64 {
	return fn2(3, 4)
}

//1、方法闭包：其实就是方法adder   返回的类型参数还是一个方法类型，和java的内部类及其类似
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}


type MyFloat float64
//2、go 没有类，但是可以为结构体定义方法（也可以为非结构体定义方法），如下：为MyFloat 定义一个方法Abs() 方法 返回 float64 类型
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//3、方法也是函数,和上一步一样
func Abs(f MyFloat) float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//4、还有一个重要的点，叫指针接收者（更常用），比如调用结构体的方法，给结构体的值重新赋值，如果用的是 值接收者，通过结构体的方法获取结构体的值时还是原来的值不会变
//如果用指针接收者则可以生效
func (v Vertex) Abs() float64  {
	return math.Sqrt(v.X*v.X+v.Y* v.Y)
}

func (v *Vertex) Score(f float64)   {//指针接收者 去掉星为值接收者
	v.X=v.X*f
	v.Y=v.Y*f
}
// 5、方法与指针重定向：有个注意的点，前面说过，方法即函数，但是在调用的时候有个区别，就是如果是函数，而且函数的参数是指针类型的，就不能传值类型的，编译会报错;
//而如果是方法，值的类型和指针的类型都可以调用这个方法。原因时GO做了一个转化，即v.Scale(5) =>(&v).Scale(5)
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}