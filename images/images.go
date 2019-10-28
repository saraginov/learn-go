package main

import (
	"fmt"
	"image"
)

/*
	Package image (https://golang.org/pkg/image/#Image) defines the package interface:

		package image
		type Image interface {
			ColorModel() color.Model
			Bounds() Rectangle
			At(x,y int) color.Color
		}

		Note: the Rectangle return value of the Bounds Method is actually an
		image.Rectangle, as the declaration is inside the package image

	The color.Color and color.Model types are also interfaces, but we'll ignore them by
	using the predefined implementations color.RGBA and color.RGBAModel.
		These interfaces and types are specified by the image/color package

*/

func imagesExample() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

func main() {
	imagesExample()
}
