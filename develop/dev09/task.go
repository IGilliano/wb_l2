package main

import (
	"flag"
	"fmt"
	"github.com/opesun/goquery"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	url := flag.String("s", "https://www.youtube.com", "url")

	flag.Parse()
	download(*url)
	parseResources(*url)
}

func download(site string) {
	resp, err := http.Get(site)
	if err != nil {
		log.Fatal(err)
	}

	defer func(resp *http.Response) {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Cant close responce body: %v", err.Error())
		}
	}(resp)

	urls := strings.Split(site, "/")
	fileName := "result/" + urls[2] + ".html"

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			fmt.Printf("Cant close file: %v", err.Error())
		}

	}(file)

	_, err = io.Copy(file, resp.Body)
}

func parseResources(site string) {
	doc, _ := goquery.ParseUrl(site)
	for _, url := range doc.Find("").Attrs("href") {
		var str []string
		switch {
		case strings.Contains(url, ".png"):
			str = strings.Split(url, "/")
			if err := downloadResources("result/"+str[len(str)-1], url); err != nil {
				fmt.Printf("Cant download from site %s:%v", url, err.Error())
			}
		case strings.Contains(url, ".jpg"):
			str = strings.Split(url, "/")
			if err := downloadResources("result/"+str[len(str)-1], url); err != nil {
				fmt.Printf("Cant download from site %s:%v", url, err.Error())
			}
		case strings.Contains(url, ".css"):
			str = strings.Split(url, "/")
			if err := downloadResources("result/"+str[len(str)-1], url); err != nil {
				fmt.Printf("Cant download from site %s:%v", url, err.Error())
			}
		}
	}
}

func downloadResources(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(resp *http.Response) {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Cant close responce body: %v", err.Error())
		}
	}(resp)

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			fmt.Printf("Cant close file: %v", err.Error())
		}
	}(file)

	_, err = io.Copy(file, resp.Body)
	return err
}
