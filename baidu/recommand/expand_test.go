package recommand

import (
	"fmt"
	"testing"
)

func TestExpandBaiduRecommendWords(t *testing.T) {
	keywords,err := Expand("大象")
	if err != nil {
		t.Errorf("get %+v\n",err)
	}
	for _, keyword := range keywords {
		fmt.Println(keyword)
	}
}
