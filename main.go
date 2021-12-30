package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/spf13/viper"
	"image/color"
	"io"
	"log"
	"strings"
	"time"
)

var config Config

func init() {
	// Load config json
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	// Set config value
	config = Config{
		EventName:      viper.GetString("config.event_name"),
		BackgroundPath: viper.GetString("config.background_img"),
		Name: Name{
			FontPath:  viper.GetString("config.name.font_path"),
			FontSize:  viper.GetFloat64("config.name.font_size"),
			PositionX: viper.GetFloat64("config.name.position_x"),
			PositionY: viper.GetFloat64("config.name.position_y"),
		},
	}

	rgbColors := strings.Split(viper.GetString("config.name.color_rgb"), ",")

	config.Name.Color = color.RGBA{
		R: StringToUint8(rgbColors[0]),
		G: StringToUint8(rgbColors[1]),
		B: StringToUint8(rgbColors[2]),
		A: 255,
	}

	// Code is optional feature
	if viper.Get("config.code") != nil {
		config.Code = Code{
			FontPath:  viper.GetString("config.code.font_path"),
			FontSize:  viper.GetFloat64("config.code.font_size"),
			PositionX: viper.GetFloat64("config.code.position_x"),
			PositionY: viper.GetFloat64("config.code.position_y"),
		}

		rgbColors := strings.Split(viper.GetString("config.code.color_rgb"), ",")

		config.Code.Color = color.RGBA{
			R: StringToUint8(rgbColors[0]),
			G: StringToUint8(rgbColors[1]),
			B: StringToUint8(rgbColors[2]),
			A: 255,
		}
	}
}

func main() {
	start := time.Now()
	fmt.Println("running app...")

	csvReader, err := NewCSVReader()
	if err != nil {
		log.Fatalln(err)
	}

	for {
		// Read csv records
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		// Generate certificate
		person := Person{
			Name: record[0],
			Code: record[1],
		}
		img, err := WriteTextOnImage(person)
		if err != nil {
			log.Fatalln(err)
		}

		err = gg.SaveJPG(fmt.Sprintf("./out/%s_%s.jpg", person.Name, config.EventName), img, 100)
		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("done, execution time =", time.Since(start))
}
