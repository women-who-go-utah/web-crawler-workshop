package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 2,
	}

	resp, err := get(client, "https://air.utah.gov")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	selection := doc.Find("a")
	if len(selection.Nodes) == 0 {
		panic("no 'a' nodes found")
	}
	for _, node := range selection.Nodes {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				fmt.Println(attr.Val)
				if strings.HasPrefix(attr.Val, "http") {
					childResp, err := get(client, attr.Val)
					if err != nil {
						fmt.Println(err)
						continue
					}
					fmt.Println(childResp.ContentLength)
				}
				break
			}
		}
	}
}

func get(client *http.Client, url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.AddCookie(&http.Cookie{
		Name:  "name",
		Value: "value",
	})
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("not a 200 status code")
	}
	return resp, err
	//defer is called here
}
