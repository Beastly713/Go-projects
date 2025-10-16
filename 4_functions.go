package main

// syntax:
// func functionName(parameter1 type, parameter2 type) returnType {
//     // function body
//     return value
// }

// add takes two integers and returns their integer sum.
// Note: (a int, b int) can be shortened to (a, b int)
// func add(a int, b int) int {
//     return a + b
// }

// func main() {
//     result := add(10, 20)
//     fmt.Println("Sum:", result) // Output: Sum: 30
// }

// IMPORTANT - multiple return values

import (
    "errors"
    "fmt"
)

// divide returns the result of a division and an error.
func divide(a, b float64) (float64, error) {
    if b == 0 {
        // Create a new error
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil // 'nil' indicates no error occurred
}

func demoDivide() {
    result, err := divide(10.0, 2.0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result) // Output: Result: 5
    }

    result, err = divide(10.0, 0.0)
    if err != nil {
        fmt.Println("Error:", err) // Output: Error: cannot divide by zero
    } else {
        fmt.Println("Result:", result)
    }
}
