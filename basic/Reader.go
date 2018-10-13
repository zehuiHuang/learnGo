package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"io"
)

//  func (T) Read(b []byte) (n int, err error)
func main()  {
var b bytes.Buffer
b.Write([] byte("首次使用"))
fmt.Fprint(&b)
b.WriteTo(os.Stdout)



//通过 string.NewReader(string) 创建一个字符串读取器
reader :=strings.NewReader("1234567890")
p := make([]byte, 4)
	for   {
		n ,err := reader.Read(p)//读取到p的缓冲区中
		if err!=nil{
			if err ==io.EOF{
				fmt.Println("EOF",n)
				break;
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(n,string(p[:n]))
	}

}

//自己实现一个 Reader
