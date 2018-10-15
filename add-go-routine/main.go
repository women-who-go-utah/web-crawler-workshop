package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	in := make(chan string, 10)
	go request(client, in)
	in <- "https://air.utah.gov"
}

func request(client *http.Client, in chan string) {
	for {
		select {
		case url := <-in:
			resp, err := get(client, url)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(resp.StatusCode)
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
	if err != nil {
		return err
	}
	defer childResp.Body.Close()
	if childResp.StatusCode != http.StatusOK {
		return fmt.Errorf("child status not 200, status: %d", childResp.StatusCode)
	}
	length := len(b)
	fmt.Println("Status Code :", childResp.StatusCode, "length:", length)
	return nil
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
}
