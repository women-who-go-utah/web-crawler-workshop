package httpclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := http.DefaultClient
	if err := get(client); err != nil {
		panic(err)
	}
}
func get(client *http.Client) error {
	resp, err := client.Get("https://air.utah.gov")
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("not a 200 status code")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
