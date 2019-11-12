package main

import (
	"fmt"
	)


type Shape2D interface{
    area() float64
    perimeter() float64
 }

type Rectangle struct{ }

func (r Rectangle) area() float64 { 
	fmt.Println("area calculated") 
	return 0.0
}
func (r Rectangle) perimeter() float64 {
	fmt.Println("perimeter calculated") 
	return 0.1
}


type Triangle struct{}
func (t Triangle) area() float64{
    fmt.Println("triangle area")
    return 0.0
}
func (t Triangle) perimeter() float64 {
   fmt.Println("triangle perimeter")
   return 0.1
}



func DrawRect( r Rectangle)    { fmt.Println("Drawing a rectangle") }
func DrawTriangle( t Triangle) { fmt.Println("Drawing a triangle")  }

func DrawShape(s Shape2D) bool {
   switch  sh := s.(type) {
   case Rectangle:
       DrawRect(sh)
   case Triangle:
       DrawTriangle(sh)
   }
   return true
}




func main() {
   var shape Shape2D
   tri := Triangle{}
   shape = tri
   DrawShape(shape)

   rec := Rectangle{}
   shape = rec
   DrawShape(rec)
   

}
