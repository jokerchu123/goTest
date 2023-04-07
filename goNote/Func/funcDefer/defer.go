package funcdefer

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func Call() {
	mu.Lock()
	mu.Unlock()
}
func CallDefer() {
	mu.Lock()
	defer mu.Unlock()
}

func Test1() {
	x, y := 1, 2
	defer func(a int) {
		fmt.Println("defer x , y = ", a, y)
	}(x)
	x += 200
	y += 100
	fmt.Println(x, y)
}
func Test2() {
	defer fmt.Println("1")
	defer fmt.Println("2")
}
func Test3() {
	test := func() (z int) {
		defer func() {
			fmt.Println("defer", z)
			z += 100
		}()
		return 100
	}
	fmt.Println("test", test())
}
