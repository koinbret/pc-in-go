package main

import (
	"fmt"
	"strings"
)

func decode(in rune) rune {
	if in >= 'a' && in <= 'z' {
		newRune := (in + 2) % 'z'
		if newRune < 'a' {
			return newRune + 'a' - 1
		}
		return newRune
	}
	return in
}

func main() {
	fmt.Println(strings.Map(decode, "g fmnc wms bgblr rpylqjyrc gr zw fylb. rfyrq ufyr amknsrcpq ypc dmp. bmgle gr gl zw fylb gq glcddgagclr ylb rfyr'q ufw rfgq rcvr gq qm jmle. sqgle qrpgle.kyicrpylq() gq pcamkkclbcb. lmu ynnjw ml rfc spj."))
	fmt.Println(strings.Map(decode, "map"))
}
