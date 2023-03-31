package a4arrayandslice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	numbers := []int{1, 2, 3, 4, 5} //数组大小被认为属于类型的一部分，不同大小的数组被认为是不同类型。因此数组显得笨重，常用切片

	got := Sum(numbers)
	want := 15

	if want != got {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) { //类型不安全，如果类型不同一样能通过编译
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTail(t *testing.T) {

	got := SumAllTail([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) { //类型不安全，如果类型不同一样能通过编译
		t.Errorf("got %v want %v", got, want)
	}
}
