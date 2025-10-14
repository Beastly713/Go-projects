package main
import "fmt"

/*
fmt.Println - Print Line
-> prints its arguments
-> adds a space between them
-> appends a newline character at the end 
*/

func demoPrintln(){
	name := "Alice"
    age := 30

    fmt.Println("Hello, World!")
    fmt.Println("Name:", name, "Age:", age) // Adds spaces automatically
    fmt.Println("---")
}

/*
fmt.Printf - Print Formatted
-> arguement : format string (text + special verbs (placehlders starting with %), followed by the values to be inserted into the placeholders)
-> you have to add the newline yourself

Formatting: You control everything using format verbs.

Common Verbs:

%v: The default format for the value. Great for general use.

%s: For strings.

%d: For decimal integers.

%f: For floating-point numbers.

%t: For booleans (true or false).

%T: To print the type of a variable.

\n: A special character for a newline.

*/

func demoPrintf() {
    name := "Bob"
    age := 25
    pi := 3.14159
    isLearning := true

    // Use verbs to format the output string. Note the \n at the end.
    fmt.Printf("User: %s is %d years old.\n", name, age)

    // Control floating-point precision (e.g., %.2f for 2 decimal places)
    fmt.Printf("The value of Pi is approximately %.2f\n", pi)

    // Print the type of a variable
    fmt.Printf("The 'age' variable has the type: %T\n", age)

    // Print a boolean value
    fmt.Printf("Is Bob learning? %t\n", isLearning)
}

/* CONTROL FLOW */

func demoIfElse() {
    score := 88

    if score > 90 {
        fmt.Println("Grade: A")
    } else if score > 75 {
        fmt.Println("Grade: B")
    } else {
        fmt.Println("Grade: C or lower")
    }
}

// if with an initialization statement:

// 'num' is only available inside this if/else block
// if num := 9; num < 0 {
//     fmt.Println(num, "is negative")
// } else if num < 10 {
//     fmt.Println(num, "has 1 digit")
// } else {
//     fmt.Println(num, "has multiple digits")
// }

// The standard "C-style" loop (3-component loop):

// Prints numbers from 0 to 4
// for i := 0; i < 5; i++ {
//     fmt.Printf("%d ", i)
// }
// Output: 0 1 2 3 4

// infinite loop
// for {
//     fmt.Println("This will run forever...")
    // You would typically have a condition inside to break out of the loop
    // if someCondition {
    //     break
    // }
// }

//  The "while" loop:
// sum := 1
// for sum < 100 { // Only the condition is present
//     sum += sum
// }
// fmt.Println(sum) // Output: 128

// The for...range loop:
// for index, char := range "Go" {
//     fmt.Printf("index: %d, character: %c\n", index, char)
// }
