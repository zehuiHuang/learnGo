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
	fmt.Println("切片并不存储数据，只是引用底层数组的一部分,如果分片改了数据，原来的数据也会改变------------------")
	names :=[4]string{"a","b","c","d"}
	fmt.Println(names)
	ab := names[0:2]
	cd :=names[1:3]
	fmt.Println(ab,cd)
	ab[1] ="f"
	fmt.Println(ab)
	fmt.Println(names)


	//数组文法,可以使用 struct
	e :=[]int{1,2,3,4}
	fmt.Println(e)
	f :=[]bool{false,true,true}
	fmt.Println(f)
	g := []struct
	{i int
	b bool}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(g)
	fmt.Println(g[1].b)

	//切片的开头和末尾都有默认值 开头默认为0,结尾默认为数组长度
	h := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(h)
	i :=h[1:]
	fmt.Println(i)
	j :=h[:2]
	fmt.Println(j)
	k :=h[:];
	fmt.Println(k)

	//切片的长度与容量,注意：如果切片设置了左边的边界，容量会在原来的基础上变化为从设置的index到原来数组长度值 之间的大小
	fmt.Println("切片的长度与容量------------------------")
	l :=[]int{1,2,3,4,5,6,7}
	printSlice(l)
	m :=l[2:]
	printSlice(m)
	n :=l[:3]
	printSlice(n)
	o :=l[1:5]
	fmt.Println(o)
	printSlice(o)

	//nil 切片,nil 切片的长度和容量为 0
	fmt.Println("nil 切片------------")
	var p []int
	fmt.Println(p)
	fmt.Println(len(p))
	if p==nil{
		fmt.Println("默认是nil")
	}


	//用 make 创建切片
	fmt.Println("用 make 创建切片--------------------------")
	q := make([]int, 5)
	printSlice2("q", q)

	r := make([]int, 0, 5)
	printSlice2("r", r)

	r2 := make([]int,5)
	printSlice2("r2", r2)

	s1 := r[:4]
	printSlice2("s1", s1)

	s2 := r[:4]
	printSlice2("s2", s2)
	t := s1[2:3]
	printSlice2("t", t)

	//向切片追加元素,在添加的时候，切片的容量并不会按照切片的大小来增加，而是按照一定规则增加
	fmt.Println("向切片追加元素-------------------------")
	var s3 []int
	printSlice(s3)

	// append works on nil slices.
	s3 = append(s3, 0)
	printSlice(s3)

	// The slice grows as needed.
	s3 = append(s3, 1)
	printSlice(s3)

	// We can add more than one element at a time.
	s3 = append(s3, 2, 3)
	//s3 = append(s3, 9)
	//s3 = append(s3, 10)
	printSlice(s3)



}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
