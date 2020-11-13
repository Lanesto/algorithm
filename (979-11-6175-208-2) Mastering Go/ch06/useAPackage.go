package main

import (
	"fmt"
	"aPackage"
)

func main() {
	fmt.Println("Using aPackage")
	aPackage.A()
	aPackage.B()
	fmt.Println(aPackage.MyConstant)
}
