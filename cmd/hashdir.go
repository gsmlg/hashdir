package main

import (
    "os"
    "fmt"
    dirhash "golang.org/x/mod/sumdb/dirhash"
)


func main() {
    str, err := dirhash.HashDir("assets", "./", dirhash.Hash1)

    if err != nil {
        fmt.Fprintf(os.Stderr, "\033[31mError: %d\033[0m", err)
        os.Exit(1)
    } else {
        fmt.Fprintf(os.Stdout, "%s", str)
        os.Exit(0)
    }
}

