package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func GetAstro() (map[string]string, error) {
	signs := []string{
		"aries",
		"taurus",
		"gemini",
		"cancer",
		"leo",
		"virgo",
		"libra",
		"scorpio",
		"sagittarius",
		"capricorn",
		"aquarius",
		"pisces",
	}

	client := NewParserHttpClient(10)

	predictions := make(map[string]string)
	for _, sign := range signs {
		url := fmt.Sprintf("https://horo.mail.ru/prediction/%s/today/", sign)
		resp, err := client.Get(url)
		if err != nil {
			return nil, err
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return nil, err
		}

		doc.Find(".article__item.article__item_alignment_left.article__item_html").Each(func(i int, s *goquery.Selection) {
			predictions[sign] = s.Text()
		})
	}
	return predictions, nil
}

func GetAstroForSign(sign string) (string, error) {
	client := NewParserHttpClient(10)

	prediction := ""
	url := fmt.Sprintf("https://horo.mail.ru/prediction/%s/today/", sign)
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	prediction = doc.Find(".article__item.article__item_alignment_left.article__item_html").Text()

	return prediction, nil
}
