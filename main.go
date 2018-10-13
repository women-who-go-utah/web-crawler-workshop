package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	//	client := http.DefaultClient
	client := &http.Client{
		Timeout: time.Second * 2,
	}

	if err := get(client); err != nil {
		panic(err)
	}
}

func get(client *http.Client) error {
	// client.Get("https://air.utah.gov")
	request, err := http.NewRequest("GET", "https://air.utah.gov", nil)
	if err != nil {
		return err
	}
	request.AddCookie(&http.Cookie{
		Name:  "name",
		Value: "value",
	})
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("not a 200 status code")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
	//defer is called here
}
