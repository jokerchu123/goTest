package typestring

import (
	"bytes"
	"strings"
	"testing"
)

// 字符串操作通常在堆上分配内存，会产生大量垃圾回收对象，建议使用[]byte缓存池
func test1() string { //该方法每次都需重新分配内存
	var s string
	for i := 0; i < 1000; i++ {
		s += "a"
	}
	return s
}
func test2() string {
	s := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		s[i] = "a"
	}
	return strings.Join(s, "")
}
func test3() string {
	var b bytes.Buffer
	b.Grow(1000)
	for i := 0; i < 1000; i++ {
		b.WriteString("a")
	}
	return b.String()
}
func BenchmarkTest1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//基准测试参数：
		// ns/op：每次迭代时间 ｜  B/op：迭代分配字节数 ｜ allocs/op：每次迭代分配内存次数

		//BenchmarkTest1-8  19327   59660 ns/op	  530276 B/op	999 allocs/op
		//test1()

		//BenchmarkTest1-8  189202	6027 ns/op	  1024 B/op	    1 allocs/op
		//test2()

		//BenchmarkTest1-8  321073	3615 ns/op	  2048 B/op	    2 allocs/op
		//test3()
	}
}
