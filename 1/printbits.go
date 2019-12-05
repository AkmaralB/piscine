package main

import (
    "os"
    "strconv"

    "github.com/01-edu/z01"
)

func main() {

    arg := os.Args
    base := "01"

    if len(arg) == 2 {
        nbr, err := strconv.Atoi(os.Args[1])
        if err != nil {
        }
        var res string
        for nbr > 0 {
            digit := base[nbr%(2)]
            res = string(digit) + res
            nbr = nbr / 2
        }

        for len(res) != 8 {
            res = "0" + res
        }

        for _, result := range res {
            z01.PrintRune(result)
        }

    }
    z01.PrintRune('\n')
