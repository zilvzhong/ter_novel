package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/modood/table"
	"strings"
	"ter_novel/config"
	"ter_novel/fetcher"
	"time"
)

func main() {
	color.Cyan(config.LOGO)
	defer func() {
		color.Yellow("期待下次使用@@")
	}()

	L:
	for {
		fmt.Fprintf(color.Output, "$%s", color.CyanString("请输入搜索的作品名称："))
		name := config.GetInputString()

		switch {
		case (len(name) == 0) :
			continue
		case name == "q":
			break L
		case name == "help":
			config.ShowHelp()
			continue
		default:
			var novel []config.Novel
			ch := make(chan config.Novel, len(config.SiteUrl))
			for i, u := range config.SiteUrl {
				go func(i int, u string) {
					fmt.Println(i, u)
					ti, addr := fetcher.Fetcher_novel(name, u)
					fmt.Println(ti,addr)
					if (ti != "" && addr != "") {
						ch <- config.Novel{i, ti, addr}
					} else {
						ch <- config.Novel{i, fmt.Sprintf("%s查询不到",name), u + name}
					}
				}(i, u)
			}

			time.Sleep(1 * 1e9)
			for _ = range config.SiteUrl {
				novel = append( novel, <- ch)
			}

			if len(novel) != 0 {
				table.Output(novel)
				P:
				for {
					fmt.Fprintf(color.Output, "$$%s", color.CyanString("请输入查看的站点Id："))
					id := strings.TrimSpace(config.GetInputString())

					switch {
					case (len(id) == 0) :
						continue
					case id == "q":
						break L
					case id == "help":
						config.ShowHelp()
						continue
					case id == "show":
						table.Output(novel)
					default:
						n, stu := config.GetNovel_id(id, novel)
						if !stu {
							fmt.Fprintf(color.Output,"$$%s", color.RedString("您输入的Id不在表Id范围，可通过show再查看表\n"))
							continue
						} else {
							chapter := fetcher.Fetcher_chapter(n.Addr)
							if len(chapter) != 0 {
								table.Output(chapter)
								for {
									fmt.Fprintf(color.Output, "$$$%s", color.CyanString("请输入查看的章节Id："))
									chapterid := strings.TrimSpace(config.GetInputString())
									switch {
									case len(chapterid) == 0 :
										continue
									case chapterid == "q":
										break L
									case chapterid == "help":
										config.ShowHelp()
										continue
									case chapterid == "return":
										break P
									default:
										c, stu := config.Getcontent_id(chapterid, chapter)
										fmt.Println(c)
										if !stu {
											fmt.Fprintf(color.Output,"$$$%s", color.RedString("您输入的Id不在表Id范围，可通过show再查看表\n"))
											continue
										} else {
											fetcher.Fetcher_content(c.Addr)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

//func main() {
//	s := "    十九年后    深秋，    冷雨连传来一阵阵的暖热，可是心中却是一片冰凉.......    为什么呢？    因"
//	fmt.Println(strings.ReplaceAll(s, "    ", "\n"))
//
//	s1 := " 十九年后    深秋，    冷雨    感觉"
//	fmt.Println(strings.ReplaceAll(s1, "    ", "\n"))
//}
