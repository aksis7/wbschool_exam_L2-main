package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"path"
	"strings"

	"golang.org/x/net/html"
)

// downloadFile скачивает и сохраняет файл
func downloadFile(fileURL string, outputPath string) error {
	resp, err := http.Get(fileURL)
	if err != nil {
		return fmt.Errorf("failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save file: %v", err)
	}

	return nil
}

// parseLinks парсит HTML и возвращает ссылки
func parseLinks(body io.Reader, baseURL string) ([]string, error) {
	doc, err := html.Parse(body)
	if err != nil {
		return nil, err
	}

	var links []string
	var crawler func(*html.Node)
	crawler = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					link := attr.Val
					u, err := url.Parse(link)
					if err != nil {
						continue
					}
					base, _ := url.Parse(baseURL)
					absURL := base.ResolveReference(u)
					links = append(links, absURL.String())
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			crawler(c)
		}
	}
	crawler(doc)

	return links, nil
}

// downloadSite рекурсивно скачивает сайт
func downloadSite(siteURL string, outputDir string, depth int) error {
	if depth == 0 {
		return nil
	}

	resp, err := http.Get(siteURL)
	if err != nil {
		return fmt.Errorf("failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()

	u, err := url.Parse(siteURL)
	if err != nil {
		return err
	}

	filePath := path.Join(outputDir, path.Base(u.Path))
	if strings.HasSuffix(siteURL, "/") || path.Ext(filePath) == "" {
		filePath = path.Join(outputDir, "index.html")
	}

	err = os.MkdirAll(path.Dir(filePath), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directories: %v", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save file: %v", err)
	}

	links, err := parseLinks(resp.Body, siteURL)
	if err != nil {
		return err
	}

	for _, link := range links {
		err := downloadSite(link, outputDir, depth-1)
		if err != nil {
			fmt.Printf("failed to download %s: %v\n", link, err)
		}
	}

	return nil
}

func main() {
	siteURL := "https://en.wikipedia.org/wiki/State_pattern"
	outputDir := "./download"
	depth := 2

	err := downloadSite(siteURL, outputDir, depth)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
