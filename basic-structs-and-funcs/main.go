package main

import "fmt"

//struct
type rectangle struct {
	length  float64
	breadth float64
	color   string
}

// funcs
func greeting(name string) string {
    fmt.Println("entering func with string: " + name)
    return "hello " + name
}

func getReference(name string) *string {
    return &name
}


func main() {
    var a string = "Robert"
    fmt.Println(greeting(a))


    b := getReference(a)

    fmt.Println(b)
    for i := 0; i < 6; i++ {
        fmt.Println(string((*b)[i]))
    }

    fmt.Println(len(a))

    var rect1 = &rectangle{10, 20, "Green"}
    fmt.Println(rect1.length)
}

