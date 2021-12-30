package main

import (
	"github.com/fogleman/gg"
	"image"
	"math"
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

	// Decrease font size 15% if name character greater than 25 characters
	fontSize := config.Name.FontSize
	if len(person.Name) > 25 {
		fontSize -= math.Round(fontSize * 0.15)
	}

	// Load font from path and set font size
	err = canvas.LoadFontFace(config.Name.FontPath, fontSize)
	if err != nil {
		return nil, err
	}

	// Set maximal width of text and color
	maxWidth := float64(imgWidth) - 60.0
	canvas.SetColor(config.Name.Color)

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

		canvas2.SetColor(config.Code.Color)
		canvas2.DrawStringWrapped(person.Code, config.Code.PositionX, config.Code.PositionY, 0.5, 0.5,
			maxWidth, 1.3, gg.AlignCenter)

		return canvas2.Image(), nil
	} else {
		return canvas.Image(), nil
	}
}
