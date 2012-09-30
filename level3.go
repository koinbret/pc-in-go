package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func FileContent(name string) []byte {
	r, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return bytes
}

func main() {
	bytes := FileContent("level3.txt")
	re, err := regexp.Compile("[^A-Z][A-Z]{3}([a-z])[A-Z]{3}[^A-Z]")
	if err != nil {
		panic(err)
	}
	matches := re.FindAllSubmatch(bytes, 1000)
	for i := 0; i < len(matches); i++ {
		match := matches[i]
		fmt.Print(string(match[1]))
	}
}
