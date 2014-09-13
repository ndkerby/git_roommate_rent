// Reading and writing files are basic tasks needed for
// many Go programs. First we'll look at some examples of
// reading files.

package main

import (
    "fmt"
    "strings"
    "io/ioutil"
)

// Reading files requires checking most calls for errors.
// This helper will streamline our error checks below.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // Perhaps the most basic file reading task is
    // slurping a file's entire contents into memory.
    dat, err := ioutil.ReadFile("../env")
    check(err)
    fmt.Printf("Start>>>\n")
    fmt.Print(string(dat))
    fmt.Printf("\n<<< end")

    var variables map[string]string
    variables = make(map[string]string)

    var rows_list = strings.Split(string(dat), "\n")
    fmt.Print(rows_list)
    for index := range rows_list {
        fmt.Printf("Row:\n")
        fmt.Print(rows_list[index])
        var stuff = strings.Split(string(rows_list[index]), "=")
        if len(stuff) == 2{
            variables[stuff[0]] = stuff[1]
        }
        fmt.Printf("\n")
    }

    fmt.Print(variables)
}
