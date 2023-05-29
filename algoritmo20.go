package main

import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    var sorted bool = true

    for i := 1; i < len(arr); i++ {
        if arr[i] < arr[i-1] {
            sorted = false
            break
        }
    }

    if sorted {
        fmt.Println("O array está em ordem crescente.")
    } else {
        fmt.Println("O array não está em ordem crescente.")
    }
}
