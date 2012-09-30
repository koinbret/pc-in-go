package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	//"regexp"
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
	bytes := FileContent("level2.txt")
	reg, err := regexp.Compile("[a-z]")
	if err != nil {
		panic(err)
	}
	s := reg.FindAll(bytes, 1000)
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[i]))
	}
}
