package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//https://www.bootschool.net/ascii  Banner在线生成工具
//http://www.network-science.de/ascii/
const LOGO = `
▄▄▄▄▄▄▄                      ▄▄   ▄  ▄▄▄▄  ▄    ▄ ▄▄▄▄▄▄ ▄     
   █     ▄▄▄    ▄ ▄▄         █▀▄  █ ▄▀  ▀▄ ▀▄  ▄▀ █      █     
   █    █▀  █   █▀  ▀        █ █▄ █ █    █  █  █  █▄▄▄▄▄ █     
   █    █▀▀▀▀   █            █  █ █ █    █  ▀▄▄▀  █      █     
   █    ▀█▄▄▀   █            █   ██  █▄▄█    ██   █▄▄▄▄▄ █▄▄▄▄▄
                     ▀▀▀▀▀▀
`

var (
	SiteUrl []string
)

type Novel struct {
	Id   int
	Name string
	Addr string
}

type Chapter struct {
	Id  int
	Name string
	Addr string
}

func init() {
	SiteUrl = []string{
		"www.dingdiann.com/searchbook.php?keyword=",
		"www.booktxt.com/search.php?q=",
		"www.qu.la/ar.php?keyWord=",
		"www.zwdu.com/search.php?keyword=",
		"www.30sy.com/search.html?searchtype=novelname&searchkey=",
	}
}

func Getenv(key, def string)  string {
	val := os.Getenv(key)
	if val == "" {
		return def
	} else {
		return val
	}
}

func GetInputString() string {
	buf := bufio.NewReader(os.Stdin)
	s, err := buf.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(s)
}

func GetNovel_id(id string, novel []Novel) (Novel, bool) {
	for _, n := range novel {
		ids, _ := strconv.Atoi(id)
		if n.Id == ids {
			return n, true
		}
	}
	return Novel{}, false
}

func Getcontent_id(id string, chapter []Chapter) (Chapter, bool) {
	for _, c := range chapter {
		ids, _ := strconv.Atoi(id)
		if c.Id == ids {
			return c, true
		}
	}
	return Chapter{}, false
}

func ShowHelp()  {
	fmt.Println()
	fmt.Println("***** 输入show查看站点源或章节")
	fmt.Println("***** 输入return时返回上一层")
	fmt.Println("***** 输入q时程序退出。")
	fmt.Println("***** 输入作品名称即可查询")
	fmt.Println("***** 输入站点源Id或章节Id即可查看内容")
}