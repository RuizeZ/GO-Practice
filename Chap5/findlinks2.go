package findLinkes2

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinkes(url)
		if err == nil {
			for _, link := range links {
				fmt.Println(link)
			}
		}
	}

}

func findLinkes(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err == nil {
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
		}
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err == nil {
			return visit(nil, doc), nil
		}
	}
	return nil, fmt.Errorf("Error")
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println(a)
				links = append(links, a.Val)
			}
		}
	}
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	return links
}
