package main
import (
    "fmt"
)
type polz struct {
    age     int
    name    string
    surname string
}
var kolich_pal int
var vtoroy_pal string
var tsifry = [] int {1,2,3,4,6,7,8,9,0}
func main() {
    Vasya := &polz{age: 32, name:"Вася", surname: "Кузнецов"}
    fmt.Print(Vasya.age)
    fmt.Print(" " + Vasya.name + " " + Vasya.surname + "\n")
    kolich_pal = 20
    fmt.Print(kolich_pal)
    fmt.Print("\n")
    blyuda := map[string]string{"cheburek": "чебурек", "shawarma" : "шаурма"}
    for m, _ := range blyuda {
        fmt.Print(blyuda[m] + " ")
    }
    vtoroy_pal = "указательный"
    fmt.Print("\n" + vtoroy_pal + "\n")
    for _, d := range tsifry {
        fmt.Print(d)
        fmt.Print(" ")
    }
}

