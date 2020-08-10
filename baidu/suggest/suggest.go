// get baidu suggest words
package suggest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type suggestJsonO struct {
	Q string `json:"q"`
	P bool   `json:"p"`
	G []struct {
		Type string `json:"type"`
		Sa   string `json:"sa"`
		Q    string `json:"q"`
	} `json:"g"`
	Slid    string `json:"slid"`
	Queryid string `json:"queryid"`
}

func MobileExpand(word string) ([]string, error) {
	eWord := escapeWord(word)

	urlFmt := "https://m.baidu.com/sugrec?pre=1&p=3&ie=utf-8&json=1&prod=wise&from=wise_web&net=1&os=1&wd=%s"
	eUrl := fmt.Sprintf(urlFmt, eWord)
	method := "GET"

	client := &http.Client{
	}
	req, err := http.NewRequest(method, eUrl, nil)

	if err != nil {
		return nil,err
	}
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Sec-Fetch-Site", "same-site")
	req.Header.Add("Sec-Fetch-Mode", "no-cors")
	req.Header.Add("Sec-Fetch-Dest", "script")
	req.Header.Add("Referer", "https://www.baidu.com/")
	req.Header.Add("Accept-Language", "en,zh-CN;q=0.9,zh;q=0.8,en-US;q=0.7,zh-TW;q=0.6")
	//req.Header.Add("Cookie", "__bsi=12192051634084907160_00_13_N_R_30_0303_c02f_Y; BDSVRBFE=Go")

	res, err := client.Do(req)
	if err != nil {
		return nil,err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil,err
	}
	d := suggestJsonO{}
	err = json.Unmarshal(body, &d)
	if err != nil {
		return nil, err
	}
	var rs []string
	for _, s := range d.G {
		rs = append(rs, s.Q)
	}
	return rs, nil
}

func PCExpand(word string) ([]string, error) {
	w := escapeWord(word)
	urlFmt := "https://www.baidu.com/sugrec?pre=1&p=3&ie=utf-8&json=1&prod=pc&from=pc_web&wd=%s&req=2&bs=%s&csor=5&pwd=%s"
	suggestUrl := fmt.Sprintf(urlFmt, w, w, w)

	method := "GET"
	client := &http.Client{
	}
	req, err := http.NewRequest(method, suggestUrl, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept", "text/javascript, application/javascript, application/ecmascript, application/x-ecmascript, */*; q=0.01")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Accept-Language", "en,zh-CN;q=0.9,zh;q=0.8,en-US;q=0.7,zh-TW;q=0.6")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	d := suggestJsonO{}
	err = json.Unmarshal(body, &d)
	if err != nil {
		return nil, err
	}
	var rs []string
	for _, s := range d.G {
		rs = append(rs, s.Q)
	}
	return rs, nil
}

func escapeWord(word string) string {
	return url.QueryEscape(word)
}
