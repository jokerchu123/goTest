package a3repeat

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	result := Repeat("a", 10)
	fmt.Println(result)
	// Output: aaaaaaaaaa
}
func TestRepeat(t *testing.T) {
	got := Repeat("a", 10)
	want := "aaaaaaaaaa"
	if got != want {
		t.Errorf("expect %v but got %v", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
