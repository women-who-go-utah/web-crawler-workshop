# Basic HTTP Client
To connect to any webservice we need an http client. Go has a default client that is sufficient for most use cases. The following instructions show how we used to default client in the web crawler.

### Setup Client
Using the Go standard library, we will implement a `DefaultClient`. First we need to import the library. Add the following code

`import "net/http"`

This imports the library. My IDE auto imports the needed libraries, so I will not be listing the imports that I make. They can be found in the acompanying [example](https://github.com/Soypete/Example-Web-Crawler/blob/master/http-client/main.go).

The client is setup in the `main()` function.

```func main() {
  client := http.DefaultClient
  if err := get(client); err != nil {
    panic(err)
  }
}
```
### Connect to Webpage

The client is passed into a `get()` function. This `get()` is used to fetched information from the site. The `net\http` package has its own `Get()` function that returns an `http.Respose` and an `error`. We can verify that the connection is good by checking the `resp.StatusCode`. To access the contents of the page we use the `ioutil` package and the `ReadAll()` funtion. This function returns a slice of bytes, so it must me cast into a string. The following is the implemented funtion.

``func get(client *http.Client) error {
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
```
