package main

import (
	"github.com/kevin-five/wordexpand/baidu/recommand"
	"log"
)

func main() {
	keywords, err := recommand.Expand("大象")
	if err != nil {
		log.Fatalf("when get baidu recommend words got: %+v\n", err)
	}
	for i, keyword := range keywords {
		log.Printf("expand word %d: %s\n", i+1, keyword)
	}
}
