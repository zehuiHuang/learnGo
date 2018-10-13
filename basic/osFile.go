package main

import (
	"os"
	"fmt"
	"io"
)


func main() {
	//proverbs := []string{
	//	"Channels orchestrate mutexes serialize\n",
	//	"Cgo is not Go\n",
	//	"Errors are values\n",
	//	"Don't panic\n",
	//}
	//file, err := os.Create("./proverbs.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//defer file.Close()
	//
	//for _, p := range proverbs {
	//	// file 类型实现了 io.Writer
	//	n, err := file.Write([]byte(p))
	//	if err != nil {
	//		fmt.Println(err)
	//		os.Exit(1)
	//	}
	//	if n != len(p) {
	//		fmt.Println("failed to write data")
	//		os.Exit(1)
	//	}
	//}
	//fmt.Println("file write done")

	//filemethod()//写入txt文件内容
	readtest()//读取txt文件内容

}

func filemethod()  {
	proverbs :=[]string{
		"aaaa","bbbbbb","cccc",
	}
	file,err :=os.Create("./text.txt")
	if err != nil{//如果没有错误说明创建text.txt成功
		fmt.Println(err)
		os.Exit(1);
	}
	defer file.Close();//defer是golang语言中的关键字，用于资源的释放，会在函数返回之前进行调用

	for _,p :=range  proverbs{
		//file 类型实现了io.Writer
		n,err :=file.Write([]byte(p)) //把数组中的其中一个字符串写到file文件中
		if err !=nil{
			fmt.Println(err)
			os.Exit(1)
		}
		if n!=len(p){//如果数组中的一个字符串的长度和写文件时返回的长度不相等说明写入的时候出错了
			fmt.Println("writer fail------")
			os.Exit(1)
		}
		fmt.Println(n)
	}
}

func  readtest()  {
	file,err := os.Open("./text.txt")
	if err !=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	p :=make([]byte,4)
	for  {
		n,err :=file.Read(p)
		if err !=nil{
			fmt.Println("读取出错了吗？？",err)//没有错，可能会返回io.EOF 表示读取完毕
		}
		if err ==io.EOF{
			break;
		}
		fmt.Println(string(p[:n]))
	}

}