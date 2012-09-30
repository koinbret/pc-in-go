package main

import (
	"archive/zip"
	"container/list"
	"fmt"
	"regexp"
)

const start = "90052"

func next(nothing string, files map[string]*zip.File) (string, string, bool) {

	file := files[nothing+".txt"]

	r, err := file.Open()
	if err != nil {
		panic(err)
	}
	defer r.Close()

	buf := make([]byte, file.UncompressedSize)
	_, _ = r.Read(buf)
	fmt.Println(string(buf))

	re, _ := regexp.Compile("nothing is (\\d+)")
	match := re.FindSubmatch(buf)
	if match != nil {
		return string(match[1]), file.Comment, true
	}
	return "", "", false
}

func main() {

	r, err := zip.OpenReader("level6_channel.zip")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	files := make(map[string]*zip.File)

	for _, f := range r.File {
		files[f.Name] = f
	}

	comments := list.New()
	cur := start
	for {
		next, comment, more := next(cur, files)
		cur = next
		if !more {
			break
		}
		comments.PushBack(comment)
	}

	for e := comments.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println("\nOXYGEN")
}
