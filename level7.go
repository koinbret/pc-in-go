package main

import (
	"fmt"
	"image/png"
	"os"
)

func main() {
	r, _ := os.Open("level7_oxygen.png")
	defer r.Close()

	img, err := png.Decode(r)
	if err != nil {
		panic(err)
	}
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	half := height / 2
	fmt.Println(width, height, half)

	for i := 0; i < width; i += 7 {
		color := img.At(i, half)
		red, _, _, _ := color.RGBA()
		fmt.Print(string(red >> 8))
	}

	res := []byte{105, 110, 116, 101, 103, 114, 105, 116, 121}
	fmt.Println("")
	fmt.Println(string(res))

}
