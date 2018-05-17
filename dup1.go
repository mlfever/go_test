package main

import (
    "os"
    "fmt"
    "bufio"
)

func main() {
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        counts[input.Text()]++
    }

    for text, num := range counts {
        if num >= 1 {
            fmt.Printf("%d %s\n", num, text)
        }
    }
}

