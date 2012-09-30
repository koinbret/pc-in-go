package main

// pickle python specific, so let's just learn how to use JSON in Go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func FileContent(name string) []byte {
	r, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer r.Close()
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return bytes
}

func main() {
	bytes := FileContent("level5_banner.json")

	var m [][][]interface{}
	json.Unmarshal(bytes, &m)

	for _, line := range m {
		for _, tuple := range line {
			how_many := int(tuple[1].(float64))
			char := tuple[0].(string)
			for i := 0; i < how_many; i++ {
				fmt.Print(char)
			}
		}
		fmt.Println("")
	}
}
