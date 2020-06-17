package main

import (
	"os"
	"io/ioutil"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
	dirhash "golang.org/x/mod/sumdb/dirhash"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func base64UrlSafeEncode(source []byte) string {
	// Base64 Url Safe is the same as Base64 but does not contain '/' and '+' (replaced by '_' and '-') and trailing '=' are removed.
	bytearr := base64.StdEncoding.EncodeToString(source)
	safeurl := strings.Replace(string(bytearr), "/", "_", -1)
	safeurl = strings.Replace(safeurl, "+", "-", -1)
	safeurl = strings.Replace(safeurl, "=", "", -1)
	return safeurl
}

func main() {
	files, err := dirhash.DirFiles("assets", "./");
	check(err)

	h := sha256.New()

	for _, f := range files {
		dat, err := ioutil.ReadFile("./assets/" + f)
		check(err)
		h.Write(dat)
	}

	ha := h.Sum(nil)
	encodHash := base64UrlSafeEncode(ha)
	fmt.Fprintf(os.Stdout, "%s", encodHash)
	
}

