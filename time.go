package main

import (
    "fmt"
    "time"
)

func main() {
    var now = time.Now()
    time.Sleep(1 * time.Second)
    var later = time.Now()
    fmt.Println(later.Sub(now))
}
