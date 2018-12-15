package main

import (
    "os"
    "fmt"
)

func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func Min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func Saturate(a, b, x int) int {
    if x < a {
        return a
    } else if x > b {
        return b
    }
    return x
}

func LogMe(fm string, args ...interface{}) {
    str := fmt.Sprintf(fm, args...)

    f, err := os.OpenFile("log", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }

    defer f.Close()

    if _, err = f.WriteString(str); err != nil {
        panic(err)
    }

    if _, err = f.WriteString("\n"); err != nil {
        panic(err)
    }
}
