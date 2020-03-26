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
// -  Package == Project == Workspace
// -  A package is a collection of common source code files
//    A package can have many related files inside of it each file ending with a file extension of go
//    The only requirement for every file inside of a package is at the very first line of each file must declare the package that it belongs to. e.g. package main
// -  There are 2 different types of packages inside the world of Go
//    1. Executable - generates a file that we can run
//    1.1 The word main is specifically used to make an executable type package
//    2. Reusable - code used as 'helpers'. Good place to put reusable logic
// -  In summary, whenever we see the word package main that means we are making an executable package. Any other name whatsoever means we're making a reusable or dependency type package
//    Anytime we make an executable package, it must always have a function inside of it called Main as well

// 3. What does 'import "fmt"' means?
// -  The import statement is used to give our package the access to some code that is written inside of another package
// -  Therefore, by saying import "fmt" specifically means give my package main access to all of the code and all the functionality that's contained inside of this other package called fmt
// -  fmt is the name of a standard library package that is included with the Go programming language by default.
//    fmt itself is a kind of abbreviated form of the word format
//    The fmt library is used to print out a lot of different information specifically to the terminal just to give you a better sense of debugging and stuff like that
// -  Besides importing a standard library like e.g. debug, fmt, math, encoding, crypto, io and etc. We can also import packages that have been authored by other engineers as well
// -  Checkout documentation around all of the standard library packages by visiting: https://golang.org/pkg/

// 4. What's that 'func' thing?
// -  func is short for function
// -  Functions inside of go function just like functions in other programming languages as well
// -  func tells go we're about to declare a function
// -  The syntax is:
//    func <sets the name of the function> (<List of arguments to pass the function>) {
//       <Function body. Calling the function runs this code>
//    }

// 5. How is the main.go file organized?
// -  In practice it always ends up being the exact same pattern:
//    At the very top we're always going to place our [package declaration]. In this case package main
//    Then right underneath that will list out all the other packages that we might need to import into this file. In this case fmt package
//    Then get down to the body of the file which is where we add in a bunch of logic that actually kind of does something. So it will be a collection of different functions, variable declarations and all that kind of other good stuff as well
