package funcdefer

import "testing"

// 性能要求高且压力大的算法应避免使用延迟调用
// BenchmarkCall-8   	150167353	         7.564 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Call()
	}
}

// BenchmarkCallDefer-8   	112667779	        10.39 ns/op	       0 B/op	       0 allocs/op
func BenchmarkCallDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallDefer()
	}
}
