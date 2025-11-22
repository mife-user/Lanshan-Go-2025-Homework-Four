package pakege

import (
	"fmt"
	"os"
	"time"
)

var cyjj string

func No_io(file *os.File) {
	fmt.Println("请输入字符串，用来测试：")
	fmt.Scanf("%s", &cyjj)
	start := time.Now()
	/*--------------------------*/
	file.WriteString(cyjj + "\n")
	/*--------------------------*/
	cost := time.Since(start)
	fmt.Printf("操作用时：%v\n", cost)
}
