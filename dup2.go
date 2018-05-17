package main

import (
    "os"
    "fmt"
    "bufio"
)

func main() {
    // def counts as a line-num map
    counts := make(map[string]int)

    // get file from args
    files := os.Args[1:]
    len := len(files)

    // if file num is zero, get file from stdin
    if len == 0 {
        countfile(os.Stdin, counts)
    } else {
        // else if files in args, set file line to count
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countfile(f, counts)
            f.Close()
        }
    }

    for line, num := range counts {
        fmt.Printf("%d\t%s\n", num, line)
    }
}

// put file line-num to count map
func countfile(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
    }
}
