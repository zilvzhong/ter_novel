package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
	"ter_novel/config"
)

func getInputString() string {
	buf := bufio.NewReader(os.Stdin)
	s, err := buf.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(s)
}

func showHelp()  {
	fmt.Println()
	fmt.Println("***** 请输入作品名称+作者，格式如：xxx作品名:作者")
	fmt.Println("***** 输入q时程序退出。")
}


func main() {
	color.Cyan(config.LOGO)
	defer func() {
		color.Yellow("期待下次使用@@")
	}()
	for {
		fmt.Fprintf(color.Output, "$$ %s", color.CyanString("请输入搜索的作品名称"))
		name := getInputString()
		if len(name) == 0 {
			continue
		}
		if name == "q" {
			break
		}
		if name == "help" {
			showHelp()
			continue
		}
	}
}
