package main

import (
	"bytes"
	"compress/bzip2"
	"fmt"
	"io/ioutil"
)

const un = "BZh91AY&SYA\xaf\x82\r\x00\x00\x01\x01\x80\x02\xc0\x02\x00 \x00!\x9ah3M\x07<]\xc9\x14\xe1BA\x06\xbe\x084"
const pw = "BZh91AY&SY\x94$|\x0e\x00\x00\x00\x81\x00\x03$ \x00!\x9ah3M\x13<]\xc9\x14\xe1BBP\x91\xf08"

func decompress(bz2 string) string {
	unbuf := bytes.NewBufferString(bz2)
	r := bzip2.NewReader(unbuf)
	bytes, _ := ioutil.ReadAll(r)
	return string(bytes)
}

func main() {
	fmt.Printf("User: %s\n", decompress(un))
	fmt.Printf("Pass: %s\n", decompress(pw))
}
