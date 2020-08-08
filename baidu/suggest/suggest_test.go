package suggest

import (
	"fmt"
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	rs, err := PCExpand("hugo")
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Join(rs, ","))
}

func TestMobileExpand(t *testing.T) {
	rs, err := MobileExpand("大象")
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Join(rs, ","))
}
