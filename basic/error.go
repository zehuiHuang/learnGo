package main

import (
	"time"
	"errors"
	"fmt"
)

type myError struct {
	when time.Time
	who  string
}

func  main()  {
	fmt.Println(validateError(-1))
	d,e :=validateError(1)
	fmt.Println(d,e)
}

func validateError(d int) (int ,error)  {
	if d <0{
		return 0,errors.New("传递的参数小于0错误")
	} else{
		return d,nil
	}
}
