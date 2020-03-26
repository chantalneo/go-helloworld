package main

import "fmt"

func main() {
	fmt.Println("Hi There!")
}

// === Five Important Questions ===

// 1. How do we run the code in our project?
// -  In the terminal, run the following command:
//    go run main.go
// -  However, GO CLI supports various commands:
//    1. go build - compiles a bunch of go source code files. If we run go build main.go, we get a main file which we can execute with ./main. To build it and then run it in the future
//    2. go run - compiles and executes one, two or a handful of files. I.e. instantly run the file
//    3. go fmt - formats all the code in each file in the current directory
//    4. go install - compiles and "installs" a package. One of the two commands that are used to handle dependencies inside of our projects
//    5. go get - downloads the raw source code of someone else's package. One of the two commands that are used to handle dependencies inside of our projects
//    6. go test - runs any tests associated with the current project

// 2. What does 'package main' mean?

// 3. What does 'import "fmt"' means?

// 4. What's that 'func' thing?

// 5. How is the main.go file organized?
