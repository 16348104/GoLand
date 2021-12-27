package main

import "fmt"

func main() {
	var a int64 = 5
	var b int64 = 8
	var c int64
	c = int64(a) + int64(b)
	//c = a+ b
	fmt.Printf("%d\n", c)
}

//第一个go程序
