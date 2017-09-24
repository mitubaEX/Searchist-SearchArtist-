package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

func main() {
	for i := 0; i <= 44; i++ {
		doc, err := goquery.NewDocument("http://www.uta-net.com/name_list/" + strconv.Itoa(i) + "/")
		if err != nil {
			fmt.Print("url scarapping failed")
		}
		doc.Find("a").Each(func(_ int, s *goquery.Selection) {
			url, _ := s.Attr("href")
			if strings.Contains(url, "artist") && len(strings.Split(url, "/")) <= 4 && strings.Split(url, "/")[3] == "" {
				// fmt.Println(url)
				fmt.Println(s.Text())
				// fmt.Println(len(strings.Split(url, "/")))
			}
		})
	}
}
