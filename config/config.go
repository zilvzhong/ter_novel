package config

import "os"

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

func Getenv(key, def string)  string {
	val := os.Getenv(key)
	if val == "" {
		return def
	} else {
		return val
	}
}

func init() {
	SiteUrl = []string{
		"www.dingdiann.com/searchbook.php?keyword=",
	}

}