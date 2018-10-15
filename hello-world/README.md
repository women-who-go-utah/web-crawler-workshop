# Hello World
It is time to start the web crawler. Using a prefered IDE make a `main.go` file. Let's start with the famous *\Hello World\* program. In the `main.go` file create a main function. It should look something like this

```
func main() {
	fmt.Println("Hello world")
}
```

Test the code. There are two ways to run the code. The first requires two lines of code.
### Build Binary and Run Executable
Go is a compiled language. The build step compiles the Go code into an executable binary. Run the following in the Command Line Terminal to build the binary. The first line accesses the repository. The second builds the binary. The third runs the executable.

```
cd hello-world
go build main.go
./main.go
```

### One Step Run
Alternatively the following code will run the Go code without building a stand alone executable.

```
cd hello-world
go run main.go
```
