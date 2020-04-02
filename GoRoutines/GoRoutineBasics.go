package main

import (
    "fmt"
    "time"
)

func hello() {
    fmt.Println("Heo world goroutine")
}
func main() {
    go hello()
    time.Sleep(1 * time.Second)
    fmt.Println("m mamta gupta kolk")
}
