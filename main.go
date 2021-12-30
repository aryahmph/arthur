package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/spf13/viper"
	"io"
	"log"
	"time"
)

var config Config

func init() {
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}

	config = Config{
		EventName:      viper.GetString("config.event_name"),
		BackgroundPath: viper.GetString("config.background_img"),
		Name: Name{
			FontPath:  viper.GetString("config.name.font_path"),
			FontSize:  viper.GetFloat64("config.name.font_size"),
			PositionX: viper.GetFloat64("config.name.position_x"),
			PositionY: viper.GetFloat64("config.name.position_y"),
		},
		Code: Code{
			FontPath:  viper.GetString("config.code.font_path"),
			FontSize:  viper.GetFloat64("config.code.font_size"),
			PositionX: viper.GetFloat64("config.code.position_x"),
			PositionY: viper.GetFloat64("config.code.position_y"),
		},
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
