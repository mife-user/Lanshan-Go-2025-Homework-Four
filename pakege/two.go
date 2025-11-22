package pakege

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var kqjj string

func Yes_io(file *os.File) {
	writeit := bufio.NewWriter(file)
	fmt.Println("请输入字符串，用来测试：")
	fmt.Scanf("%s", &kqjj)
	start := time.Now()
	/*--------------------------*/
	writeit.WriteString(kqjj)
	writeit.Flush()
	/*--------------------------*/
	cost := time.Since(start)
	fmt.Printf("操作用时：%v\n", cost)
}
