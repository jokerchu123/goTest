package main

import (
	"context"
	"fmt"
	"goTest/Dict/dict"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
example: simpleDict hello
		`)
		os.Exit(1)
	}
	word := os.Args[1]
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//go dict.QueryDict1(ctx, cancel, word)
	go dict.QueryDict2(ctx, cancel, word)
	select {
	case <-ctx.Done():
		return
	}
}
