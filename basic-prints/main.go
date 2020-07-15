package main

import "fmt"

const BUFFER_SIZE = int(5)

func main() {
    fmt.Println("Hello, World!")

    // var keyword
    var name = "Robert"
    var age int = 26

    // pointer stuff in go
    var pAge *int = nil
    var np *string = &name

    fmt.Println(name, age)
    fmt.Println(np)
    fmt.Println(*np)

    pAge = new(int)
    *pAge = 1
    for i:= 1; i < BUFFER_SIZE; i++ {
        *pAge++
    }
    fmt.Println(*pAge)


    var a[BUFFER_SIZE]int
    fmt.Println("empt: ", a)
    for i:= 1; i < BUFFER_SIZE; i++ {
        a[i] = i
    }
    fmt.Println(a)
}

