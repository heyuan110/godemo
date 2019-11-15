/*
@File    :   main.go
@Time    :   2019/01/25 14:02:33
@Author  :   Bruce
@Version :   1.0
@Contact :   bruce.he@patpat.com
@License :   (C)Copyright 2019, patpat.com
@Desc    :   None
*/

package main

import (
	"fmt"
	"test/testlib"
)

func main() {
	fmt.Println("hello")
	fmt.Println("1+2=", testlib.Add(1, 2))
	fmt.Println("3-2=", testlib.JianNum(3, 2))
}
