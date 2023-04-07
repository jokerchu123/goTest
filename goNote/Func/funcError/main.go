package main

import (
	"fmt"
	"log"
)

type DivError struct {
	x, y int
}

func (DivError) Error() string { //实现error接口的Error方法，DivError成为一个error类型
	return "division by zero"
}

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, DivError{x, y}
	}
	return x / y, nil
}
func main() {
	z, err := div(5, 0)
	if err != nil {
		switch e := err.(type) {
		case DivError:
			fmt.Println(e, e.x, e.y)
		default:
			fmt.Println(err)
		}
		log.Fatalln(err) //Fatal系列函数会在写入日志信息后调用os.Exit(1);Panic系列函数会在写入日志信息后panic

	}
	fmt.Println(z)

}
