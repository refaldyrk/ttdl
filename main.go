package main

import (
	"fmt"

	"github.com/refaldyrk/ttdl/service"
)

func main() {
	yt := service.New()

	result := yt.Get("https://www.tiktok.com/@yucholoco/video/7268539324663188741").GetLink()
	for _, v := range result {
		fmt.Println(v)
	}
}
