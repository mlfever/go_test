package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    counts := make(map[string]int)
    files := os.Args[1:]
    for _, arg := range files {
        readline, err := ioutil.ReadFile(arg)
        if (err != nil) {
            fmt.Fprintf(os.Stderr, "%v", err)
            continue
        }

        for _, line:= range strings.Split(string(readline), "\n") {
            counts[line]++
            //fmt.Printf("%s\n", line)
        }

        for line, num := range counts {
            fmt.Printf("%d\t%s\n", num, line)
        }
    }
}

