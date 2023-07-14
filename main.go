package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.OnResponse(func(r *colly.Response) {

		u, err := url.Parse(r.Request.URL.String())
		if err != nil {
			fmt.Println("Gagal memperoleh nama file:", err)
			return
		}
		fileName := strings.TrimPrefix(u.Hostname(), "www.") + ".html"
		err = os.WriteFile(fileName, r.Body, 0644)
		if err != nil {
			fmt.Println("Gagal menyimpan hasil crawl:", err)
			return
		}

		fmt.Println("Hasil crawl berhasil disimpan dalam file:", fileName)
	})

	fmt.Print("Masukkan URL yang ingin di-crawl: https://")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	url := scanner.Text()

	err := c.Visit("https://" + url)
	if err != nil {
		fmt.Println("Gagal melakukan crawling:", err)
		return
	}
}
