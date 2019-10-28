package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image is "custom" image package "Image" type implementation
type Image struct {
	/*
		I think image should implement PicTwo
		and return what Pic in exercise-slices returns or something along those lines,
		instead of a dx, dy guess Pic would be a method to this?

		Hypothesis above incorrect, initial dx, dy int assumption was correct

	*/
	d1x, d1y, d2x, d2y int
}

/*
	From docs (https://golang.org/pkg/image/#Image)
	type Image interface {
		ColorModel()
		// returns the Image's color model.
		Bounds() Rectangle
		// Bounds returns the domain for which At can return non-zero color;
		// Bounds does not necessarily contain the point (0,0)
		At(x, y int) color.Color
		// At returns the color of the pixel at (x,y)
		// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid
		// At(Bounds()Max.X-1, Bounds().Max.Y-1) returns the lower-right one
	}
*/

// ColorModel is an Image interface method implementation
func (img Image) ColorModel() color.Model {
	// ColorModel should return color.RGBAModel.
	// colorModel() Method must take a Rectangle type image,
	// which is why Bounds() must return a rectangle

	// ColorModel implements the Image interface.
	return color.RGBAModel
}

// Bounds is an Image interface method implementation
func (img Image) Bounds() image.Rectangle {
	// Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).

	// Image is a pointer in the src files (https://golang.org/src/image/image.go?s=14584:14618#L496),
	// should we implement it as a pointer here or no?

	// From docs (https://golang.org/pkg/image/#Rectangle)
	// Rect is shorthand for Rectangle{Pt(x0, y0), Pt(x1, y1)}.
	// The returned rectangle has minimum and maximum coordinates swapped if necessary so that it is well-formed.

	// THEREFORE: I do not think we should implement the pointer,
	// because we don't need to modify "m" in main()

	// return image.NewRGBA(image.Rect(0, 0, 100, 100)) // for reference

	// example return, since image.NewRGBA != image.Rectangle type
	// return image.Rect(0, 0, 100, 100)
	return image.Rect(img.d1x, img.d1y, img.d2x, img.d2y)
}

// At is an Image interface method implementation
func (img Image) At(x, y int) color.Color {
	// At should return a color;
	// the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.

	// with slice we get 255 blue value, 255(ie full opacity) and we have our dx dy for the v values

	// return color.RGBA(img.dx, img.dy)  -> did not work because color.RGBA does not have default values

	// After looking at package documentation, these are the arguments we need to pass to RGBA()
	// https://golang.org/pkg/image/color/#RGBA
	// therefore implementing
	/*
		type RGBA struct {
			R, G, B, A uint8
		}q
	*/

	// for example:
	// return RGBA(x, y, 255, 255) -> this does not work because arg(s) must be cast as type(s) uint32
	// return RGBA(uint8(x), uint8(y), uint8(255), uint8(255))

	// blue := math.Sqrt(float64(x + y))
	blue := x ^ 2 + y ^ 2
	// return color.RGBA{uint8(x), uint8(y), 255, 255} // original
	return color.RGBA{uint8(x), uint8(y), uint8(blue), 255}

}

func main() {
	m := Image{50, 30, 100, 200}
	pic.ShowImage(m)
}
