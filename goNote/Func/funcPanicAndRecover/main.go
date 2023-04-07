package main

import (
	"fmt"
	"log"
)

//除非是不可恢复性，导致系统无法正常工作的错误，否则不建议使用panic
//panic和recover是go的内置函数，panic会立即中断当前函数流程，执行延迟调用。在延迟调用函数中，recover会捕获并返回panic提交的错误对象

func test1() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(err)
		}
	}()
	panic("dead!")
	fmt.Println("exit!") //不会到达
}
func test2() { //连续调用panic，仅向最后一个会被recover捕获
	defer func() {
		for {
			if err := recover(); err != nil {
				log.Println(err)
			} else {
				log.Fatalln("fatal!")
			}
		}
	}()
	defer func() {
		panic("You are dead!")
	}()
	panic("I am dead!")
}
func main() {
	//test1()
	test2()
}
