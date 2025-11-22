package main

import (
	"Lanshan-homework/four/pakege"
	"fmt"
	"os"
)

var chooseit int = 0

func main() {
	File, err1 := os.Create("D:/gogogo/go/Lanshan-Go-2025-Homework/four/test.txt")
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("选择方式（0-无缓冲，1-有缓冲)")
	fmt.Scanf("%d", &chooseit)
	if chooseit == 0 {
		pakege.No_io(File)
	} else if chooseit == 1 {
		pakege.Yes_io(File)
	} else {
		fmt.Println("还是头文件好用")
	}
	err2 := File.Close()
	if err2 != nil {
		return
	}

}
