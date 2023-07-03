## Setting up GO Environment and Hello World Program

### Setting up GO Environment

After we set up our environment and verify it,
we’ll build a simple program, learn about the different ways to build and run Go code,
and then explore some tools and techniques that make Go development easier.

### Installing GO
To write Go code, you first need to download and install the Go development tools.
The latest version of the tools can be found at the downloads page on the [Go website](https://go.dev/).
Choose the download for your platform and install it.

I am using Linux so I'll be downloading the Linux version of the Go tools.

### Go installation on Linux
``` bash
1. Download the archive and extract it into /usr/local, creating a Go tree in /usr/local/go. For example:
$ wget https://golang.org/dl/go1.17.linux-amd64.tar.gz

2. Extract the archive into /usr/local, creating a Go tree in /usr/local/go. For example:
$ sudo tar -C /usr/local -xzf go1.17.linux-amd64.tar.gz

3. Add /usr/local/go/bin to the PATH environment variable. You can do this by adding this line to your /etc/profile (for a system-wide installation) or $HOME/.profile:
$ export PATH=$PATH:/usr/local/go/bin

4. Verify that you've installed Go by opening a command prompt and typing the following command:
$ go version
```

### Hello World Program

Now that we have installed Go, we can write our first Go program.

Create a file named hello.go and add the following code to it:
``` go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

# Understanding the code

- The first line of the program package main defines the package name in which this program will reside.
  Every Go program must start with a package declaration.

- The next line import "fmt" imports the fmt package into our program.
  The fmt package implements formatted I/O functions similar to C's printf and scanf.

- The next line func main() is the main function where the program execution begins.
  The main function is the entry point to our program.

- The next line fmt.Println(...) is another function available in the fmt package that prints the        message to the console.


### Go Tools

Go comes with a set of command-line tools that you can use to format, build, and run your code.
We’ll look at some of the most common tools in this section.

### go run and go build

To run the program, put the code in hello.go and use go run.
``` bash
$ go run hello.go
```
Note: Use go run when you want to treat a Go program like a script and
run the source code immediately.

To build a binary, put the code in hello.go and use go build.
``` bash
$ go build hello.go
```
Note: Use go build when you want to build a binary that you can distribute and run later. 
In the above case the binary file is [hello_world](https://github.com/satyampsoni/golang/blob/master/helloworld/hello_world)

