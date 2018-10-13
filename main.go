package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func main() {
	//	client := http.DefaultClient
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	resp, err := get(client, "https://air.utah.gov")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	//	if len(doc.Nodes) == 0 {
	//		panic("no nodes found")
	//	}
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
				if strings.HasPrefix(attr.Val, "http") {
					err = getChild(client, attr)
					if err != nil {
						fmt.Println(err)
					}
				}
				break
			}
		}
	}
}

func getChild(client *http.Client, attr html.Attribute) error {
	fmt.Println(attr.Val)
	childResp, err := get(client, attr.Val)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(childResp.Body)
	defer childResp.Body.Close()
	if err != nil {
		return err
	}
	length := len(b)
	fmt.Println("Status Code :", childResp.StatusCode, "length:", length)
	return nil
}

func get(client *http.Client, url string) (*http.Response, error) {
	// client.Get("https://air.utah.gov")
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
	//	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("not a 200 status code")
	}
	//	body, err := ioutil.ReadAll(resp.Body)
	//	if err != nil {
	//		return err
	//	}
	//	fmt.Println(string(body))
	return resp, err
	//defer is called here
}
