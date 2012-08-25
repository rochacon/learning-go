package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    m := make(map[string]int)

    words := strings.Fields(s)

    for _, word := range words {
        if m[word] != 0 {
            m[word]++
        } else{
            m[word] = 1
        }
    }

    return m
}

func main() {
    wc.Test(WordCount)
    print("\n")
}
