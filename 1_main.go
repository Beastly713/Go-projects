package main

import "fmt"

func main() {
    fmt.Println("Hello from Go!")
    demonstrate()
    demoPrintln()
    demoPrintf()
    demoIfElse()
    demoDivide()
}

/*
1. source file -> text file that contains human-readable instructions written in a specific programming language

compilation : source file -> machine-readable executable file(series of binary ones and zeroes)


2. package <name> is a Declaration
When you write package main, you are telling the compiler: "The code in this file belongs to a special package that should be compiled into a runnable, executable program."

example:

Imagine your file system looks like this:

/my-go-projects
└── /hello
    └── hello.go
The Directory: The directory for your main package is /my-go-projects/hello. The moment you put a file with package main inside it, that directory becomes the location of your main package for this specific program.

The Source File(s): The source file is hello.go. If you were to add another file, say goodbye.go, into that same /hello directory and put package main at the top of it, both hello.go and goodbye.go would be the source files for your single main package.

in my case
Go-projects -> directory for main package
main.go -> source files
any other .go file -> source files

What is special about the main package?

-> A package is a collection of source files in the same directory that are compiled together.

why compiled together?
If it's a main package, all the .go files are compiled and linked together to create one single executable file.

If it's a library package (like our mymath example), they are compiled into one single package archive (a .a file). This archive is a pre-compiled, reusable block of code that other programs can then import and use.

NOTE : main is a special package name. When the Go compiler sees package main, it knows that this package is intended to be an executable program (one you can run directly from your terminal), rather than a library (a collection of code to be used by other programs).

Behavior 1: package main
Keyword: main

Intent: "This is an executable program."

Compiler's Action: When the compiler sees package main, it knows its job is to produce a standalone, runnable application (an .exe file on Windows, or a binary file on Linux/macOS).

ehavior 2: Library Package (e.g., package mymath, package database)
Keyword: Any valid package name that is not main.

Intent: "This is a collection of reusable code (a library)."

Compiler's Action: The compiler knows this code isn't meant to be run by itself. Its job is to produce a package archive (a .a file). This file is a pre-compiled, efficient bundle of code that other packages (including a main package) can import and use.

example:

/myproject
└── /mymath
    ├── add.go
    └── subtract.go
add.go file:

Go

package mymath

// Add returns the sum of two integers.
func Add(a, b int) int {
    return a + b
}
subtract.go file:

Go

package mymath

// Subtract returns the difference of two integers.
func Subtract(a, b int) int {
    return a - b
}
Both files start with package mymath and are in the mymath directory. They are part of one single package.

*/

/*-------------------------------------------------------------------------------------------------------------------*/

/*

1. import : include code from other packages in your program

2. "fmt": This is the name of the package being imported. fmt (short for "format") is a standard Go library package that contains functions for formatted I/O (Input/Output).

3. func : keyword used to declare a function

4. main() : special function name -> When you run an executable program built with Go, the program starts by executing the code inside the main function. It is the entry point of your application. Every executable Go program must have a main package and a main function.

5. .fmt: This specifies that we are using a function from the fmt package that we imported earlier.

*/

