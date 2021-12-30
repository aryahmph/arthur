package main

import (
	"github.com/fogleman/gg"
	"image"
	"image/color"
)

// WriteTextOnImage write text on image based on Person data, it will return image.Image and error.
func WriteTextOnImage(person Person) (image.Image, error) {
	// Load Image
	loadImage, err := gg.LoadImage(config.BackgroundPath)
	if err != nil {
		return nil, err
	}

	// Get width and height from loaded image
	imgWidth, imgHeight := loadImage.Bounds().Dx(), loadImage.Bounds().Dy()

	// Create new canvas and put loaded image on canvas
	canvas := gg.NewContext(imgWidth, imgHeight)
	canvas.DrawImage(loadImage, 0, 0)

	// Load font from path and set font size
	err = canvas.LoadFontFace(config.Name.FontPath, config.Name.FontSize)
	if err != nil {
		return nil, err
	}

	// Set maximal width of text and color
	maxWidth := float64(imgWidth) - 60.0
	canvas.SetColor(color.Black)

	// Write text on image
	canvas.DrawStringWrapped(person.Name, config.Name.PositionX, config.Name.PositionY, 0.5, 0.5,
		maxWidth, 1.3, gg.AlignCenter)

	// if config for code is set, then write text again
	if config.Code.FontPath != "" {
		canvas2 := gg.NewContext(imgWidth, imgHeight)
		canvas2.DrawImage(canvas.Image(), 0, 0)

		err = canvas2.LoadFontFace(config.Code.FontPath, config.Code.FontSize)
		if err != nil {
			return nil, err
		}

		canvas2.SetColor(color.Black)
		canvas2.DrawStringWrapped(person.Code, config.Code.PositionX, config.Code.PositionY, 0.5, 0.5,
			maxWidth, 1.3, gg.AlignCenter)

		return canvas2.Image(), nil
	} else {
		return canvas.Image(), nil
	}
}
