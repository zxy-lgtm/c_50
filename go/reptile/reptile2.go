package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	var keyword string
	fmt.Println("请输入你需要搜索的关键词:")
	fmt.Scanln(&keyword)

	num := 1

	var number int
	fmt.Println("请输入你需要几页的内容:")
	fmt.Scanln(&number)

	for i := 1; i <= number; i++ {
		requestUr1 := "https://so.csdn.net/so/search/s.do?q=" + keyword + "&t=all&platform=pc&p=" + strconv.Itoa(i) + "&s=&tm=&v=&l=&lv=&u=&ft="
		rp, err1 := http.Get(requestUr1)
		if err1 != nil {
			panic(err1)
		}
		body, err2 := ioutil.ReadAll(rp.Body)
		if err2 != nil {
			panic(err2)
		}

		content := string(body)
		defer rp.Body.Close()

		//fmt.Println(content)

		dom, err3 := goquery.NewDocumentFromReader(strings.NewReader(content))
		if err3 != nil {
			panic(err3)
		}
		dom.Find(".search-list-con").Each(func(i int, selection *goquery.Selection) {
			//fmt.Println(selection.Text())
			selection.Find(".datetime").Each(func(i int, title *goquery.Selection) {
				fmt.Printf("%-3d", num)
				fmt.Println(title.Text())
				num++
			})
		})
	}

}
