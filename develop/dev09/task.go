package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go-wget <url>")
		return
	}

	startURL := os.Args[1]
	u, err := url.Parse(startURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Неправильный URL: %v\n", err)
		return
	}

	visited := make(map[string]bool)
	err = downloadPage(u, u.Host, visited)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка скачивание страницы: %v\n", err)
	}
}

func downloadPage(u *url.URL, baseDir string, visited map[string]bool) error {
	if visited[u.String()] {
		return nil
	}
	visited[u.String()] = true

	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Создание локального пути для сохранения файла
	localPath := filepath.Join(baseDir, u.Path)
	if strings.HasSuffix(localPath, "/") {
		localPath = filepath.Join(localPath, "index.html")
	} else if filepath.Ext(localPath) == "" {
		localPath += ".html"
	}

	// Создание директорий
	if err := os.MkdirAll(filepath.Dir(localPath), 0755); err != nil {
		return err
	}

	// Сохранение файла
	file, err := os.Create(localPath)
	if err != nil {
		return err
	}
	defer file.Close()

	var doc *html.Node
	if strings.HasSuffix(localPath, ".html") {
		doc, err = html.Parse(resp.Body)
		if err != nil {
			return err
		}
		// Анализ HTML и загрузка ресурсов
		err = downloadResources(doc, u, baseDir, visited)
		if err != nil {
			return err
		}
		html.Render(file, doc)
	} else {
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Downloaded %s\n", u.String())
	return nil
}

func downloadResources(n *html.Node, baseURL *url.URL, baseDir string, visited map[string]bool) error {
	if n.Type == html.ElementNode {
		var attrName string
		switch n.Data {
		case "img", "script":
			attrName = "src"
		case "link":
			if val := getAttr(n, "rel"); val != "stylesheet" {
				return nil
			}
			attrName = "href"
		case "a":
			attrName = "href"
		}

		if attrName != "" {
			for i, attr := range n.Attr {
				if attr.Key == attrName {
					resourceURL, err := baseURL.Parse(attr.Val)
					if err != nil {
						return err
					}
					if resourceURL.Host != baseURL.Host {
						return nil
					}
					downloadPage(resourceURL, baseDir, visited)
					localPath := filepath.Join(baseDir, resourceURL.Path)
					n.Attr[i].Val, _ = filepath.Rel(filepath.Dir(filepath.Join(baseDir, baseURL.Path)), localPath)
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if err := downloadResources(c, baseURL, baseDir, visited); err != nil {
			return err
		}
	}

	return nil
}

func getAttr(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}
