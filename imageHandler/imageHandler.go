package imageHandler

import (
	"image/color"
	"image/png"
	"os"
)

var RGBArray []color.RGBA

func SetRGBArrayFromFile(file *os.File) {
	img, err := png.Decode(file)
	if err != nil {
		panic(err)
		return
	}

	// Iterate over every pixel in the image
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Get the color of the pixel
			c := img.At(x, y)

			// Convert it to RGBA
			r, g, b, a := c.RGBA()

			// Normalize to 8-bit values
			r8 := uint8(r >> 8)
			g8 := uint8(g >> 8)
			b8 := uint8(b >> 8)
			a8 := uint8(a >> 8) // Alpha value, you can ignore if not needed

			// Append the color to the array
			RGBArray = append(RGBArray, color.RGBA{R: r8, G: g8, B: b8, A: a8})
		}
	}
}
