package main

import (
    "fmt"
    "github.com/melomilight/go-homeworks/shapes"
)

func main() {
    circle := shapes.NewCircle(5.0)
    rectangle := shapes.NewRectangle(4.0, 6.0)

    shapesList := []shapes.Shape{circle, rectangle}

    for _, shape := range shapesList {
        fmt.Printf("%s - Area: %.2f, Perimeter: %.2f\n", shape.GetName(), shape.Area(), shape.Perimeter())
    }
}