package Oublie_Go

import "fmt"

func Oublie(name string) {
	if name == "XYQ" || name == "ZYH" {
		fmt.Println("太阳的光还耀眼吗")
	} else {
		fmt.Println("你好，我在等我曾经的太阳")
	}
}

func V2Test() {
	fmt.Println("你用的是v2版本，对应v2分支")
}
