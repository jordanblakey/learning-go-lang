package main

import "fmt"

// Comments

/*
Multi Line Comments
*/

func main() {

  fmt.Println("Hello World!")

////////////////////////////////////
// DATA TYPES                     //
////////////////////////////////////

  // Note that data types are static (cannot be changed once declared)

  // uint8
  // uint16
  // uint32
  // uint64

  // int8
  // int16
  // int32
  // int64

  // float32
  // float64


  ////////////////////////////////////////////////////
  // VARIABLES
  ////////////////////////////////////////////////////

  var age int = 10
  var floatingvar float64 = 1.12340339

  randNum := 1 // Type inferencing

  fmt.Println(floatingvar)
  fmt.Println(age, floatingvar, randNum) // spaces added

  var numOne = 1.000
  var num99 = .9999

  fmt.Println(numOne - num99)


  ////////////////////////////////////////////////////
  // ARITHMETIC OPERATORS +-*/%
  ////////////////////////////////////////////////////

  fmt.Println("6 + 4 =", 6 + 4)
  fmt.Println("6 - 4 =", 6 - 4)
  fmt.Println("6 * 4 =", 6 * 4)
  fmt.Println("6 / 4 =", 6 / 4)
  fmt.Println("6 % 4 =", 6 % 4)


  ////////////////////////////////////////////////////
  // DIFFERENT WAYS TO DECLARE AND PRINT VARIABLES
  ////////////////////////////////////////////////////
  const pi float64 = 3.14159265

  var (
    // varA = 2
    // varB = 3
    myName = "Jordan"
    isOver40 bool = false
  )

  fmt.Println(len(myName))
  fmt.Println(myName + " is over 40? ", isOver40)

  fmt.Println(myName + " is a robot \n \n")
  fmt.Println("New lines above!")


  // Printf provides more ways to format output, but doesn't end with a newline and so these must be added manually.
  fmt.Printf("%f \n", pi) // Print a float
  fmt.Printf("%.3f \n", pi) // Print a float to a specified decimal point precision
  fmt.Printf("%T \n", pi) // Print the data type
  fmt.Printf("%t \n", isOver40) // Print a boolean
  fmt.Printf("%d \n", 100) // Print an integer
  fmt.Printf("%b \n", 100) // Print binary representation of the second argument
  fmt.Printf("%x \n", 100) // Print hex representation of the second argument
  fmt.Printf("%e \n", pi) // Print scientific notation representation of the second argument

  ////////////////////////////////////////////////////
  // LOGICAL OPERATORS
  ////////////////////////////////////////////////////

  



  // END FUNC MAIN
}
