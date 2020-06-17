package main

import (
	"os"
	"io"
	"io/ioutil"
	"bytes"
	"fmt"
	dirhash "golang.org/x/mod/sumdb/dirhash"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func open(p string) (io.ReadCloser, error) {
	dat, err := ioutil.ReadFile("./assets/" + p)
	check(err)

	return ioutil.NopCloser(bytes.NewReader(dat)), nil
}

func main() {
	files, err := dirhash.DirFiles("assets", "./");
	check(err)

	str, err := dirhash.Hash1(files, open)
	check(err)

	fmt.Fprintf(os.Stdout, "%s", str)
}

