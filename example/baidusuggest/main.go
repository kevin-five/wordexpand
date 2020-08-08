package main

import (
	"github.com/kevin-five/wordexpand/baidu/suggest"
	"log"
)

func main() {
	keywords, err := suggest.PCExpand("大象")
	if err != nil {
		log.Fatalf("when get baidu suggest words got: %+v\n", err)
	}
	for i, keyword := range keywords {
		log.Printf("expand word %d: %s\n", i+1, keyword)
	}
}
