// 字符串索引与字符串截取
package main

import (
	"fmt"
	"strings"
)

func main() {
	n := "hello-world-world-你好呀"
	m := strings.Index(n, "world")
	fmt.Println("获取子字符串 world 的最开始位置：", m)
	l := strings.LastIndex(n, "world")
	fmt.Println("获取子字符串 world 在最末端的位置：", l)

	k := n[m:] // 字符串可作为字符组成的数组：可索引、可迭代
	fmt.Println("截取 m 往后的字符串：", k)
	p := n[m : m+3]
	fmt.Println("截取 m 往后的 3 位字符串：", p)
}
