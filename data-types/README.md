## Setting Up Setting Up Your Go Environment and Hello World Program

### Installing GO tools
- To write Go code, you first need to download and install the Go development tools.
  The latest version of the tools can be found at the downloads page on the Go website.

- [Download Go](https://golang.org/dl/), Choose the download for your platform and install it

- I am on linux so I will be downloading the linux version of the Go tools. I will be using the command line to install the Go tools. I will be using the wget command to download the Go tools. I will be using the following command to download the Go tools.

  ```bash

// Downloading the Go tools
  1. wget https://dl.google.com/go/go1.14.2.linux-amd64.tar.gz

// Extracting the Go tools
  2. tar -xvf go1.14.2.linux-amd64.tar.gz

// Moving the Go tools to the /usr/local directory
  3. sudo mv go /usr/local

// Setting up the Go environment
  4. nano ~/.bashrc

// Adding the following lines to the .bashrc file
5. export GOROOT=/usr/local/go
    export GOPATH=$HOME/go
    export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

// Sourcing the .bashrc file
6. source ~/.bashrc

// Checking the Go version
7. go version
  ```

Out of the box, Go ships with many development tools. You access these tools via the
go command. They include a compiler, code formatter, linter, dependency manager,
test runner, and more

### Hello World Program
- To write a hello world program in Go, create a file called hello.go and add the following code to it.

  ```go
  package main

  import "fmt"

  func main() {
    fmt.Println("Hello World")

  }

  #### Running the Hello World Program

  ```
- initialize a module for your project, use the following command.

  ```bash
    go mod init filename.go
  ```
- To run the hello.go program, use the following command.

  ```bash
    go run hello.go
    ```
- To build the hello.go program, use the following command.
    
      ```bash
     go build hello.go
     ```
## Explantion of the Hello World Program
- The first line of the program package main defines the package name. Every Go program must start with a package declaration. Packages are Go's way of organizing and reusing code. There are two types of packages in Go: executable and reusable. Executable packages create an executable file that we can run. Reusable packages are used as libraries in other programs. In our example, we are creating an executable package called main. Therefore, we use package main as the first line of code in all our example programs.

- The second line import "fmt" tells Go that we want to use the fmt package from the standard library. The fmt package implements formatting for input and output.The fmt package implements the Writer interface and exports various functions to format text. These functions implement operations for formatting (writing) and scanning (reading) text. We will learn more about interfaces in the next chapter.

- The third line func main() is the main function where the program execution begins. The main function should always be present in your program. Importance of main function will be explained in the next chapter. The opening curly brace { should always be in the same line as the function declaration. If you use a new line, Go will throw an error. you can read more about [Go's coding style guidelines](https://golang.org/doc/effective_go.html#formatting).

- The fourth line fmt.Println("Hello World") is used to print "Hello World" to the console. The fmt.Println() function is used to print text to the console. The text to be printed should be enclosed inside double quotes (" ").






 