package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
)

func GetTitles() {
	resp, err := http.Get("https://pyxt.ustc.edu.cn/?menu=public_replay")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	res := make([][]string, 0)
	res = append(res, []string{"id", "title", "name"})

	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		row := s.Find("td")
		if row.Length() == 9 && strings.Contains(row.Eq(2).Text(), "软件工程") {
			res = append(res, []string{trim(row.Eq(0).Text()), trim(row.Eq(1).Text()), trim(row.Eq(3).Text())})
		}
	})

	fd, err := os.OpenFile("titles.csv", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	c := csv.NewWriter(fd)

	for _, r := range res {
		if err := c.Write(r); err != nil {
			log.Fatal(err)
		}
	}
	c.Flush()
	if err := c.Error(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Press Enter to exit...")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}

func trim(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		return unicode.IsSpace(r)
	})
}
