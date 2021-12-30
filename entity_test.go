package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")

	err := viper.ReadInConfig()
	if err != nil {
		t.Fatal(err)
	}

	loadedConfig := Config{
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

	fmt.Printf("%+v\n", loadedConfig)

	assert.NotEmpty(t, loadedConfig.EventName)
	assert.NotEmpty(t, loadedConfig.BackgroundPath)
	assert.NotEmpty(t, loadedConfig.Name)
	assert.NotEmpty(t, loadedConfig.Code)
}

func TestLoadConfigOptionalCode(t *testing.T) {
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.SetConfigName("app.config")

	err := viper.ReadInConfig()
	if err != nil {
		t.Fatal(err)
	}

	if viper.Get("config.code") == nil {
		fmt.Println("config code is optional")
	}

	loadedConfig := Config{
		EventName:      viper.GetString("config.event_name"),
		BackgroundPath: viper.GetString("config.background_img"),
		Name: Name{
			FontPath:  viper.GetString("config.name.font_path"),
			FontSize:  viper.GetFloat64("config.name.font_size"),
			PositionX: viper.GetFloat64("config.name.position_x"),
			PositionY: viper.GetFloat64("config.name.position_y"),
		},
	}

	fmt.Printf("%+v\n", loadedConfig)

	assert.NotEmpty(t, loadedConfig.EventName)
	assert.NotEmpty(t, loadedConfig.BackgroundPath)
	assert.NotEmpty(t, loadedConfig.Name)
}
