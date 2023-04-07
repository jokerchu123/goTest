package main

import (
	"fmt"
	"sync"
	"time"
)

func test1(num int) {
	var wg sync.WaitGroup
	var sendRPC = func(x int) {
		fmt.Println(x)
	}
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(x int) {
			sendRPC(x)
			wg.Done()
		}(i)
		//wrong use,i被覆盖
		// go func() {
		// 	sendRPC(x)
		// 	wg.Done()
		// }()
	}
	wg.Wait()

}
func test2() {
	//var mu sync.Mutex
	done := make(chan bool, 1)
	periodic := func() {
		//defer mu.Unlock()
		for {
			fmt.Println("tick")
			time.Sleep(1 * time.Second)
			//mu.Lock()
			if len(done) != 0 && <-done {
				return
			}
			//mu.Unlock()
		}
	}
	time.Sleep(1 * time.Second)
	fmt.Println("started")
	go periodic()
	time.Sleep(5 * time.Second)
	//mu.Lock()
	done <- true
	//mu.Unlock()
	fmt.Println("cancelled")
	time.Sleep(3 * time.Second)
}
func test3() {
	counter := 0
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			counter += 1
			mu.Unlock()
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(counter)
}
func test4() { //原子操作
	alice := 1000
	bob := 1000
	var mu sync.Mutex
	total := alice + bob
	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			alice++
			//mu.Unlock()
			//mu.Lock()
			bob--
			mu.Unlock()
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			bob++
			//mu.Unlock()
			//mu.Lock()
			alice--
			mu.Unlock()
		}
	}()
	start := time.Now()
	for time.Since(start) < 1*time.Second {
		mu.Lock()
		if alice+bob != total {
			fmt.Printf("unequal! alice+bob = %v,total = %v", alice+bob, total)
		}
		mu.Unlock()
	}
	fmt.Println(alice, bob)
}
func main() {
	//test1(100)
	//test2()
	//test3
	test4() //原子操作测试
}
