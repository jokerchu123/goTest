#  **https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/hello-world**

# 测试
## 单元测试
>go test -v

    func ExampleFunc() {
	result := Func()
	fmt.Println(result)
	// Output: wantResult
    }

    func TestFunc(t *testing.T) {
	got := Func()
	want := wantResult
	if got != want {
		t.Errorf("expect %v but got %v", want, got)
	}
}
## 基准测试
> go test -bench=.    

    func BenchmarkFunc(b *testing.B) {
        for i := 0; i < b.N; i++ {
            Func()
        }
    }
## 覆盖率测试
> go test -cover
# 文档
    go install golang.org/x/tools/cmd/godoc
    godoc -http :8000  
    localhost:8000/pkg