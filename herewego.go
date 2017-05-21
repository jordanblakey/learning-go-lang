package main

import "fmt"

// Comments

/*
Multi Line Comments
*/

func main() {

  fmt.Println("Hello World!")

////////////////////////////////////////////////////
// DATA TYPES                     //
////////////////////////////////////////////////////

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

  fmt.Println("true && false =", true && false)
  fmt.Println("true || false =", true || false)
  fmt.Println("!true =", !true)


////////////////////////////////////////////////////
// COMPARISON/RELATIONAL OPERATORS
////////////////////////////////////////////////////

  // == != <= >= < >

////////////////////////////////////////////////////
// LOOPS
////////////////////////////////////////////////////

  // For loops
  i := 1

  for i <= 10 {
    fmt.Println(i)

    i++ // i = i + 1
    // i-- // i = i - 1

  }

  // For loops shorthand
  for j := 0; j < 5; j++ {
    fmt.Println(j)
  }

////////////////////////////////////////////////////
// IF STATEMENTS
////////////////////////////////////////////////////

  yourAge := 18

  if yourAge >= 16 {
    fmt.Println("You can drive.")
  } else if yourAge >= 18 {
    fmt.Println("You can drive, and vote. Woohoo!")
  } else {
    fmt.Println("You can't drive.")
  }


////////////////////////////////////////////////////
// SWITCH STATEMENTS
////////////////////////////////////////////////////

  switch yourAge {
    case 16: fmt.Println("Go drive!")
    case 18: fmt.Println("Go vote!")
    default: fmt.Println("Go have fun!")
  }

////////////////////////////////////////////////////
// ARRAYS
////////////////////////////////////////////////////

  var favNums2[5] float64
  favNums2[0] = 163
  favNums2[1] = 78557
  favNums2[2] = 691
  favNums2[3] = 3.141
  favNums2[4] = 1.618

  fmt.Println(favNums2[3])

  // Alternative array syntax

////////////////////////////////////////////////////
// FOR EACH LOOPS
////////////////////////////////////////////////////

  favNums3 := [5]float64 {1, 2, 3, 4, 5}

  for i, value := range favNums3 {
    fmt.Println(value, i)
  }

  for _, value := range favNums3 { // Using and underscore for the index variable loops through with no access to the index number of the value
    fmt.Println(value)
  }

////////////////////////////////////////////////////
// SLICES
////////////////////////////////////////////////////

  // Create a new arrray by slicing from an existing array

  numSlice := []int {5, 4, 3, 2, 1}
  numSlice2 := numSlice[3:5]
  numSlice3 := make([]int, 5, 10)
  copy(numSlice3, numSlice)

  fmt.Println("numSlice2[0] =", numSlice2[0])
  fmt.Println("numSlice2[1] =", numSlice2[1])
  fmt.Println("numSlice[:2] =", numSlice[:2])
  fmt.Println("numSlice[2:] =", numSlice[2:])
  fmt.Println("numSlice3[0:] =", numSlice3[0:])

  numSlice3 = append(numSlice3, 0, -1)
  fmt.Println("numSlice3[6] =", numSlice3[6])

////////////////////////////////////////////////////
// MAPS
////////////////////////////////////////////////////

  presAge := make(map[string] int)
  presAge["TheodoreRoosevelt"] = 42
  presAge["John F. Kennedy"] = 43

  fmt.Println("presAge length:", len(presAge))
  fmt.Println(presAge["TheodoreRoosevelt"])

  delete(presAge, "John F. Kennedy")
  fmt.Println("presAge length:",len(presAge))

////////////////////////////////////////////////////
// CALLING FUNCTIONS
////////////////////////////////////////////////////

  // SIMPLE FUNCTION ADDING ALL ARRAY ELEMENTS
  listNums := []float64{1,2,3,4,5}
  fmt.Println("Sum :", addThemUp(listNums))

  // STORE MULTIPLE RETURNED VALUES INTO VARS
  num1, num2 := next2Values(5)
  fmt.Println(num1, num2)

  // UNDEFINED NUMBER OF ARGUMENTS
  fmt.Println(subtractThem(1,2,3,4,5))

  // RECURSIVE FUNCTION (FACTORIAL)
    fmt.Println(factorial(3))
    // 3 * factorial(2) == 3 * 2 = 6
    // 2 * factorial(1) == 2 * 1 = 2
    // factorial(0) == 1

  // DEFER FUNCTIONS
  defer printTwo()
  printOne()


  fmt.Println("safeDiv(3,0):", safeDiv(3,0))
  fmt.Println("safeDiv(3,2):", safeDiv(3,2))

  demPanic()

  // Pointers
  x := 0
  changeXVal(x)
  fmt.Println("x =", x)

  y := 0
  changeYValNow(&y)
  fmt.Println("y =", y)

////////////////////////////////////////////////////
// NESTED FUNCTIONS VIA ClOSURES
////////////////////////////////////////////////////

  num3 := 3
  doubleNum := func() int {

    num3 *= 2

    return num3
  }

  fmt.Println(doubleNum())
  fmt.Println(doubleNum())


////////////////////////////////////////////////////
// END MAIN FUNCTION: DECLARATIVE ONLY AFTER THIS POINT
////////////////////////////////////////////////////
}


////////////////////////////////////////////////////
// CREATING FUNCTIONS
////////////////////////////////////////////////////

func addThemUp(numbers []float64) float64 { // array type, return type
  sum := 0.0

  for _, val := range numbers {
    sum += val // sum = sum + val
  }

  return sum
}

func next2Values(number int) (int, int) {
  return number+1, number+2
}

func subtractThem(args ...int) int {
  finalValue := 0

  for _, value := range args {
    finalValue -= value
  }

  return finalValue
}


////////////////////////////////////////////////////
// RECURSIVE FUNCTION
////////////////////////////////////////////////////

func factorial(num int) int {
  if num == 0 {
    return 1
  }

  return num * factorial(num - 1)
}

////////////////////////////////////////////////////
// DEFER FUNCTION
////////////////////////////////////////////////////

func printOne(){fmt.Println(1)}
func printTwo(){fmt.Println(2)}


////////////////////////////////////////////////////
// SAFE FUNCTIONS
////////////////////////////////////////////////////

func safeDiv(num1, num2 int) int{
  defer func() {
    fmt.Println(recover())
  }()

  solution := num1 / num2
  return solution

}

////////////////////////////////////////////////////
// PANIC
////////////////////////////////////////////////////

  func demPanic(){
    defer func(){
      fmt.Println(recover())
    }()

    panic("PANIC")
  }

////////////////////////////////////////////////////
// POINTERS
////////////////////////////////////////////////////

  func changeXVal(x int) {
    x = 2
  }

  func changeYValNow(y *int) {
    *y = 2
  }
