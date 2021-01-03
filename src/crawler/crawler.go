//a concurrent web crawler who explore  the link graph of the web in breadth-first order
//To test, go run crawler.go https://www.catie.fr
package main

import (
	"fmt"
	"net/http"
    "os"
	"log"
	"golang.org/x/net/html"
)

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	//TODO
	res, error := http.Get(url)
    if error != nil {
        log.Print(error)
    }
    doc, error := html.Parse(res.Body)
    if error != nil {
        log.Print(error)
    }
    var docLinks []string
    visitNode := func(node *html.Node) {
        if node.Type == html.ElementNode && node.Data == "a" {
            for _, data := range node.Attr {
                if data.Key == "href" {
                    docLinks = append(docLinks, data.Val)
                    break
                }
            }
        }
    }
	forEachNode(doc, visitNode, nil)
	return docLinks, nil
}

//!-Extract

// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		//recursive call
		//TODO
		forEachNode(c, pre, nil)
	}
	if post != nil {
		post(n)
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(list)
	return list
}

//!+
func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	//unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		//TODO
		crawl((<- worklist) [0])
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	//TODO
}
