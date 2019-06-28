package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strconv"
    "github.com/PuerkitoBio/goquery"
)

func Scrape(w string) {
    query := fmt.Sprintf("https://eow.alc.co.jp/search?q=%s", w)
    res, err := http.Get(query)
    if err != nil {
        log.Fatal(err)
    }
    doc, err := goquery.NewDocumentFromReader(res.Body);

    //html, err := doc.Find("body").Html()
    //fmt.Print(html)

    num, _ := strconv.Atoi(doc.Find("#itemsNumber > strong").Text())
    if num == 0 {
        os.Exit(1)
    } 

    word := doc.Find("#resultsList > ul:nth-child(3) > li:nth-child(1) > span:nth-child(1) > h2:nth-child(1) > span:nth-child(1)").Text()
    fmt.Print(word)
    fmt.Print("\n----\n")
    doc.Find("#resultsList > ul:nth-child(3) > li:nth-child(1) > div:nth-child(2) > ol:nth-child(2) > li").Each(func(idx int, s *goquery.Selection) {
             idx++
             fmt.Printf("%d\n", idx)
             fmt.Printf("%s\n\n", s.Text())
         })
}

func main() {
    if len(os.Args) != 2 {
        os.Exit(1)
    }
    //fmt.Printf("args: %s", os.Args[1])
    Scrape(os.Args[1])
}
