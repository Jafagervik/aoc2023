package main

import (
    "fmt"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    dat, err := os.ReadFile("./day1.txt")
    check(err)
    fmt.Print(string(dat))

    for i, v := range dat {
        fmt.Println(i, v)
    }

}
