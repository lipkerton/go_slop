package main

import (
    "os"
    "encoding/json"
    "fmt"
)

type Entity struct {
    Name    string  `json:"name"`
    Surname string  `json:"surname"`
    Age     int     `json:"age"`
}

func main() {
    var e Entity
    var f *os.File
    var err error
    
    f, err = os.OpenFile("credentials.json", os.O_RDONLY, 0111)
    if err != nil {
        fmt.Println(err)
        return
    }
    if err := json.NewDecoder(f).Decode(&e); err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(e)
}
