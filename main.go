package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	//	client := http.DefaultClient
	client := &http.Client{
		Timeout: time.Second * 2,
	}

	resp, err := get(client)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	goquery.NewDocumentFromReader(resp.Body)
}

func get(client *http.Client) (*http.Response, error) {
	// client.Get("https://air.utah.gov")
	request, err := http.NewRequest("GET", "https://air.utah.gov", nil)
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
	return nil, err
	//defer is called here
}
