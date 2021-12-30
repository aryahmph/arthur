package main

import "image/color"

type Config struct {
	EventName      string
	BackgroundPath string
	Name           Name
	Code           Code
}

type Name struct {
	FontPath  string
	FontSize  float64
	PositionX float64
	PositionY float64
	Color     color.RGBA
}

type Code struct {
	FontPath  string
	FontSize  float64
	PositionX float64
	PositionY float64
	Color     color.RGBA
}

type Person struct {
	Name string
	Code string
}
