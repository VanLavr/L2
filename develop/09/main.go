package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type parser struct {
	pageID int
	depth  int
}

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	var (
		url   string
		depth int
	)
	flag.StringVar(&url, "u", "", "provide url for downloading")
	flag.IntVar(&depth, "d", 1, "provide depth for downloading referred pages")
	flag.Parse()

	if url == "" {
		fmt.Println("provide a link to download")
		return
	}

	if depth < 1 {
		fmt.Println("provided depth is invalid")
		return
	}

	p := new(parser)
	p.depth = depth
	p.parseURL(url)
}

func (p *parser) parseURL(url string) {
	p.pageID++
	if p.pageID > p.depth {
		return
	}

	if err := os.Mkdir("./download", os.ModePerm); err != nil {
		if err.Error() != "mkdir ./download: file exists" {
			log.Fatal(err)
		}
	}

	file, err := os.Create(fmt.Sprintf("./download/file%d", p.pageID))
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}(file)

	c := colly.NewCollector(colly.ParseHTTPErrorResponse())

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("status", r.StatusCode)
	})

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		p.parseLinks(h)
	})

	c.OnResponse(func(r *colly.Response) {
		file.Write(r.Body)
	})

	c.Visit(url)
}

func (p *parser) parseLinks(links *colly.HTMLElement) {
	for _, link := range links.DOM.Nodes {
		for _, attr := range link.Attr {
			if len(attr.Val) < 10 {
				continue
			} else if attr.Val[:8] == "https://" {
				p.parseURL(attr.Val)
			}
		}
	}
}
