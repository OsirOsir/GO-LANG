package main

import (
	"fmt"
	"image/color"

	"github.com/fogleman/gg"
)

func main() {
	dc := gg.NewContext(200, 100)
	dc.SetRGB(0, 0, 0)
	dc.Clear()

	// Draw badge background
	dc.SetRGB(0.8, 0.2, 0.2)
	dc.DrawCircle(100, 50, 40)
	dc.Fill()

	// Draw text
	dc.SetColor(color.White)
	dc.DrawStringAnchored("Badge", 100, 50, 0.5, 0.5)

	// Save to PNG file
	dc.SavePNG("badge.png")
	fmt.Println("Badge generated successfully!")
}
