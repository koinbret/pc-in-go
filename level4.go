package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const urlTemplate = "http://www.pythonchallenge.com/pc/def/linkedlist.php?nothing=%s"

const start = "12345"

func next(nothing string) []byte {
	url := fmt.Sprintf(urlTemplate, nothing)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
	return body
}

func main() {
	re, err := regexp.Compile("next nothing is (\\d+)")
	if err != nil {
		panic(err)
	}

	nothing := start
	for {
		body := next(nothing)
		sbody := string(body)
		if strings.HasSuffix(sbody, ".html") {
			break
		}

		if strings.HasPrefix(sbody, "Yes. Divide by two and keep going.") {
			i, err := strconv.Atoi(nothing)
			if err != nil {
				panic(err)
			}
			nothing = fmt.Sprintf("%d", i/2)
			continue
		}
		nothing = string(re.FindSubmatch(body)[1])
	}
}
