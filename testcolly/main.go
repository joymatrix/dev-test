package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0"),
	)
	c.Async = true

	// 添加图像爬虫规则
	c.OnHTML(".cimg", func(e *colly.HTMLElement) {
		fmt.Println("===========", e.Attr("src"))
		// 下载图像
		imgurl := e.Attr("src")

		if imgurl != "" && strings.Contains(imgurl, "https") {
			imgurls := strings.Split(imgurl, "?")
			if len(imgurls) > 0 {
				e.Request.Visit(imgurls[0])
			}

		}

	})

	c.OnResponse(func(r *colly.Response) {

		fmt.Println("filename:", r.Request.URL, r.FileName())
		fileName := r.FileName()
		names := strings.Split(fileName, ".")
		if len(names) > 0 && names[1] != "unknown" {
			newFilename := names[1] + ".jpg"
			f, err := os.Create(newFilename)
			if err != nil {
				fmt.Println("Error creating file:", err)
				return
			}
			defer f.Close()

			_, err = f.Write(r.Body)
			if err != nil {
				fmt.Println("Error writing file:", err)
				return
			}

		}

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	parkName := url.QueryEscape("深圳市东湖公园")
	visitUrl := fmt.Sprintf("https://cn.bing.com/images/async?q=%s&first=1&count=30&mmasync=1", parkName)
	// 开始爬取
	c.Visit(visitUrl)
	c.Wait()
}
