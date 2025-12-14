package main

import (
	"errors"
	"fmt"
	"log"
)

func add(f int64, s int64) (int64, error) {
    if f < 0 && s < 0 {
        return 0, errors.New("cannot add signed numbers")
    }
    return f + s, nil
}

func main() {
    result, err := add(-10, -20)
    if err != nil {
        log.Print(err)
    }
    fmt.Println(result)
}
