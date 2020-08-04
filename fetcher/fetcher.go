package fetcher

import (
	"bufio"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"net/url"
	"os"
	"strings"
	"ter_novel/config"
)

type Course struct {
	Title       string
	Description string
	Creator     string
	Level       string
	URL         string
	Language    string
	Commitment  string
	HowToPass   string
	Rating      string
}

//func Newfetcher(name, uri string)  {
//	c := colly.NewCollector(
//		colly.CacheDir("./coursera_cache"),
//		)
//	cc := colly.NewCollector()
//	q := url.QueryEscape(name)
//	ul := "http://" + uri + q
//	//detailCollector := c.Clone()
//	//courses := make([]Course, 0, 200)
//
//	extensions.RandomUserAgent(c)
//	extensions.Referer(c)
//
//	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
//		link := e.Attr("href")
//		if !(strings.TrimSpace(e.Text) == name) {
//			return
//		}
//		fmt.Println(link, strings.TrimSpace(e.Text))
//		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
//		cc.Visit(e.Request.AbsoluteURL(link))
//	})
//
//	cc.OnHTML("dl", func(e *colly.HTMLElement) {
//		e.DOM.Find("a[href]").Each(func(i int, selection *goquery.Selection) {
//			fmt.Println(selection.Attr("href"))
//			fmt.Println(selection.Text())
//		})
//	})
//
//
//	c.OnRequest(func(request *colly.Request) {
//		fmt.Println("Visiting", request.URL.String())
//	})
//
//	//c.OnResponse(func(r *colly.Response) {
//	//	host := r.Request.URL.Host
//	//	fmt.Println(host)
//	//	fmt.Println(string(r.Body))
//	//})
//	c.Visit(ul)
//}

func Fetcher_novel(name, uri string) (ti, address string) {
	c := colly.NewCollector()
	q := url.QueryEscape(name)
	ul := "http://" + uri + q

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if !(strings.TrimSpace(e.Text) == name) {
			return
		}
		bufReader := bufio.NewReader(strings.NewReader(string((e.Response.Body))))
		doc, _ := goquery.NewDocumentFromReader(bufReader)
		ti = strings.TrimSpace(doc.Find("title").Text())
		address = e.Request.AbsoluteURL(link)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError", err)
		return
	})
	c.Visit(ul)
	c.Wait()
	return ti, address
}

func Fetcher_chapter(uri string) []config.Chapter {
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	extensions.Referer(c)
	var chapter []config.Chapter

	c.OnHTML("dl", func(e *colly.HTMLElement)  {
		e.DOM.Find("a[href]").Each(func(i int, selection *goquery.Selection) {
			link, _:= selection.Attr("href")
			addr := e.Request.AbsoluteURL(link)
			chapter = append(chapter,config.Chapter{i, selection.Text(), addr})
		})
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError", err)
	})
	c.Visit(uri)
	c.Wait()
	return chapter
}

func Fetcher_content(uri string)  {
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError", err)
	})

	c.OnHTML("div#content", func(e *colly.HTMLElement) {
		data := strings.ReplaceAll(e.Text, "    ", "\n")
		f, err := os.OpenFile("./data.txt", os.O_APPEND, 0660)
		if err != nil {
			f, _ = os.Create("./data.txt")
		}
		f.WriteString(data)
		//fmt.Println(strings.ReplaceAll(e.Text, "    ", "\n"))
	})
	c.Visit(uri)
}
