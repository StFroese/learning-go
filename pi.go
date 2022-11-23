package main

import (
    "fmt"
    "math"
    "math/rand"
)

func main() {
    inside := 0
    for i := 1; i < 10000000; i++ {
        x := rand.Float64()
        y := rand.Float64()
        if math.Pow(x,2) + math.Pow(y,2) <= 1 {
            inside += 1
        }
    }
    ratio := float64(inside)/float64(10000000)
    fmt.Println(4*ratio)
}
