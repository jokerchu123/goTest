package main

import funcdefer "goTest/goNote/Func/funcDefer"

func main() {
	//参数值在调用时被复制并保存
	//funcdefer.Test1()

	//多个延迟调用注册按FILO次序执行
	//funcdefer.Test2()

	//return 和 panic语句都会终止当前函数，引发延迟调用，return语句不同于ret汇编指令，他会优先更新返回值
	funcdefer.Test3()
}
