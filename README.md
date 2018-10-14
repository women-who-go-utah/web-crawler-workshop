# Women Who Go Workshop Oct 13th, 2018
## Intro to Go: Web Crawler
This workshop was put on by Zach Badgett. This is a step-by-step walk through walktrough of what was covered at the worshop.

## Install Go
Follow this link to install the Go binary on your computer. [download](https://golang.org/dl/) Follow the instructions to setup you go path as an environment variable.

## Make a Repo
The next step is to create a directory. Following go convention we will add this folder to the GOPATH.

`mkdir -p $GOPATH/src/github.com/soypete/web-crawler`

This makes the all of the directories in the path.

## Start Project
It is time to start the web crawler. Using a prefered IDE make a `main.go` file. Let's start with the famous *\Hello World\* program. In the `main.go` file create a main function. It should look something like this

```
func main() {
	fmt.Println("Hello world")
}
```
