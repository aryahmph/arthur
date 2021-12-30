package main

import (
	"log"
	"strconv"
)

func StringToUint8(numberStr string) uint8 {
	numberInt, err := strconv.Atoi(numberStr)
	if err != nil {
		log.Fatalln(err)
	}
	return uint8(numberInt)
}
