/*
Go is a statically-typed language, which means you must declare the type of a variable before you can use it.

Data types
- int(and others)
- float32, float64
- bool
- string
*/

/*-------------------------------------------------------------------------------------------------------------------*/


/*
variable declaration
#
var keyword
eg: var age int 

NOTE: If a variable is declared but not given an initial value, it is automatically assigned its "zero value".
0 for all integer types
0.0 for all float types
false for booleans
"" (an empty string) for strings

#
var with an Initializer
eg : var age int = 30

#
type inference 
// Go infers that 'age' is an 'int' because 30 is an integer.
var age = 30

// Go infers that 'name' is a 'string'.
var name = "Alice"

#
shorthand method via :=  operator
eg:
func someFunction() {
    // Declares and initializes 'age' as an int with the value 30.
    age := 30
    
    // Declares and initializes 'name' as a string.
    name := "Alice"
    
    // Declares and initializes 'isLearning' as a bool.
    isLearning := true
}

- for changing the value further use =
- Important: The := operator can only be used inside a function.

*/

/*-------------------------------------------------------------------------------------------------------------------*/

/*
Constants

- declared using const keyword
- must be initialised
- value must be known at compile time -> you can't assign the value of a function call to a const 
eg: // const CompileTime = time.Now() // <-- This will cause a compile error.

*/

package main

import "fmt"

// A constant declared at the package level (outside of a function).
const WelcomeMessage = "Learning Go!"

func demonstrate() {
    // --- CONSTANTS ---
    const Pi = 3.14

    // --- VARIABLES ---
    
    // Using the short declaration operator `:=`
    language := "Go"
    version := 1.18 // Go infers this is a float64
    isAwesome := true
    
    // Using the 'var' keyword for a variable we don't initialize yet.
    // It will get its "zero value", which is 0 for an int.
    var year int
    
    // --- PRINTING VALUES AND TYPES ---
    // We use fmt.Printf with '%v' to print the value and '%T' to print the type.
    fmt.Println(WelcomeMessage)
    
    fmt.Printf("Language: %v (type: %T)\n", language, language)
    fmt.Printf("Version: %v (type: %T)\n", version, version)
    fmt.Printf("Is it awesome? %v (type: %T)\n", isAwesome, isAwesome)
    fmt.Printf("The year is %v (zero value, type: %T)\n", year, year)
    
    // --- MODIFYING A VARIABLE ---
    year = 2025 // Assign a new value to 'year'.
    fmt.Printf("The year is now %v\n", year)
    
    // --- USING A CONSTANT ---
    var radius = 5.0
    // The calculation is done using the constant 'Pi'.
    circumference := 2 * Pi * radius
    fmt.Printf("A circle with radius %v has a circumference of %v\n", radius, circumference)
}

// Learning Go!
// Language: Go (type: string)
// Version: 1.18 (type: float64)
// Is it awesome? true (type: bool)
// The year is 0 (zero value, type: int)
// The year is now 2025
// A circle with radius 5 has a circumference of 31.400000000000002
