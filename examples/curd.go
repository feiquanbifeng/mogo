package main

import (
    "fmt"
    "mogo"
)

type People struct {
    Name string
    Age  int64
    Sex  string
}

func main() {
    // mogo := &Mogo{}
    mogo := mogo.Mogon()
    mogo.Connect("mongodb://localhost/test")

    /*People{
        "TT",
        27,
        "N",
    }*/

    c := mogo.Model("people", new(People))
    /*switch mode.(type) {
      case *People:
          fmt.Println("peoe;")
      default:
      }*/
    fmt.Println(c)

    err := c.Insert(&People{"Ale", 23, "+55 53 8116 9639"},
        &People{"Cla", 45, "+55 53 8402 8510"})
    if err != nil {
        panic(err)
    }
}
