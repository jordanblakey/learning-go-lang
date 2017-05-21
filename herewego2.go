package main

import "fmt"
import "math"

func main() {

  // rect1 := Rectangle{ leftX: 0, topY:50, height:10, width:10}
  rect2 := Rectangle{ 0, 50, 10, 10}

  fmt.Println("Rectangle is", rect2.width, "wide.")
  fmt.Println("Area of rectangle =", rect2.area())

  rect := Rectangle{0, 0, 20, 50}
  circ := Circle{4}

  fmt.Println("Rectangle area is =", rect.area())
  fmt.Println("Circle area is =", circ.area())

}
/// END FUNC MAIN ///


type Shape interface {
area() float64
}

type Rectangle struct {

  leftX float64
  topY float64
  height float64
  width float64

}

type Circle struct {
  radius float64
}

func (r Rectangle) area() float64{
  return r.height * r.width
}

func (c Circle) area() float64{
  return math.Pi * math.Pow(c.radius, 2)
}

func getArea(shape Shape) float64{
  return shape.area()
}

// func(rect *Rectangle) area() float64{
//   return rect.width * rect.height
// }
