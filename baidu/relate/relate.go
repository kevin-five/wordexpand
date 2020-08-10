package relate

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/kevin-zx/baidu-seo-tool/search"
	"strings"
)

func PCExpand(keyword string) ([]string, error) {
	html, err := search.GetBaiduPCSearchHtmlWithRN(keyword, 1, 10)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}
	var rs []string
	doc.Find("#rs th a").Each(func(_ int, selection *goquery.Selection) {
		rs = append(rs, selection.Text())
	})
	return rs, nil
}

func MobileExpand(keyword string) ([]string,error) {
	html,err := search.GetBaiduMobileSearchHtml(keyword,1)
	if err != nil {
		return nil,err
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}
	var rs []string
	doc.Find("#page-relative .rw-list-new a span").Each(func(_ int, selection *goquery.Selection) {
		rs = append(rs, selection.Text())
	})
	return rs, nil
}
