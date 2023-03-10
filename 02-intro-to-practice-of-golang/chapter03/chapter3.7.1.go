// 双引号与反引号在字符串赋值中的区别
package main

import (
	"fmt"
)

func main() {
	var a string
	var b string

	a = "hello world\nGo"
	// 双引号的字符串赋值：支持转义，不支持字面量打印与多行打印。
	b = `hello
	     \n
		 world`
	// 反引号的字符串赋值：不支持转义，支持字面量打印与多行打印。

	fmt.Printf("双引号的字符串：%v\n", a)
	fmt.Printf("反引号的字符串：%v\n", b)
}

/*
   变量的内存分配说明：
   Golang 中的字符串变量声明与赋值，其变量名称实际为指向变量值的指针，
   变量名称不存储变量的值。当程序运行时，程序向系统申请一块连续的内存
   区域用于存储变量值，比如在 64 位体系架构的机器上，一个 16 进制的
   虚拟内存地址对应的空间存储 8 字节数据，因此，若实际的字符串变量的值
   可能超出 8 字节数据，将使用一段连续内存来存储，而对变量的指针取值操作
   返回的是起始的内存地址。
*/