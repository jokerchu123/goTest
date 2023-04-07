package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func string_Basic() {
	//*1
	s := `line\r\n,
	line 2`
	fmt.Println(s) //``支持定义不做转义处理的原始字符串，支持跨行

	//*2
	s0 := "abcdefg"
	s1 := s0[:3]
	// StringHeader于string头结构相同，unsafe。Pointer用于指针类型转换
	//打印结果二者相同，即返回子串时，底层依旧指向原字节数组
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s0))) //&reflect.StringHeader{Data:0x100344267, Len:7}
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s1))) //&reflect.StringHeader{Data:0x100344267, Len:7}

	//*3
	//修改字符串时需将其转换为可变的[]byte或[]rune类型，完成后再转换为string类型，需要进行复制操作有内存开销
	//某些时候可采用非安全方法进行数据转换
	toString := func(bs []byte) string {
		return *(*string)((unsafe.Pointer(&bs)))
	}
	bs := []byte("hello,world")
	ss := toString(bs)
	//%x取地址
	fmt.Printf("%x\n", bs) //68656c6c6f2c776f726c64
	fmt.Printf("%x\n", ss) //68656c6c6f2c776f726c64

	//*4
	//UTF-8编码条件下，一个汉字占3个字节
	su := "你好！"
	//utf8.RuneCountInString()获取准确unicode字符数量
	fmt.Println(len(su), utf8.RuneCountInString(su)) //9 ,3
}
func main() {
	string_Basic()
}
