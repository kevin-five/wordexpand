package relate

import (
	"log"
	"testing"
)

func TestPCExpand(t *testing.T) {
	keyword := "大象"
	rs,err := PCExpand(keyword)
	if err != nil {
		t.Log(err)
	}
	for _, r := range rs {
		log.Printf("%s\n",r)
	}

}

func TestMobileExpand(t *testing.T) {
	keyword := "大象"
	rs,err := MobileExpand(keyword)
	if err != nil {
		t.Log(err)
	}
	for _, r := range rs {
		log.Printf("%s\n",r)
	}
}
