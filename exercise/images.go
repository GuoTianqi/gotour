package exercise

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 400, 400)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{
		uint8((x + y) % 255),
		uint8((x + y*2) % 255),
		uint8((x*2 + y) % 255),
		uint8((x + y*4) % 255),
	}
}

func ImagesMain() {
	m := Image{}
	pic.ShowImage(m)
}
