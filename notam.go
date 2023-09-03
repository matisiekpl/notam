package notam

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Fetch(icao string) ([]string, error) {
	data := url.Values{
		"reportType":    {"Report"},
		"actionType":    {"notamRetrievalByICAOs"},
		"submit":        {"View NOTAMs"},
		"retrieveLocId": {icao},
	}
	response, err := http.PostForm("https://www.notams.faa.gov/dinsQueryWeb/queryRetrievalMapAction.do", data)
	if err != nil {
		return nil, fmt.Errorf("cannot load faa resource: %w", err)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("cannot parse faa response: %w", err)
	}
	if strings.Contains(string(body), "Location not covered") || strings.Contains(string(body), "Your session has expired or an invalid request has occurred.") {
		return nil, errors.New("cannot fetch notam for invalid icao code")
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		return nil, fmt.Errorf("cannot parse faa response: %w", err)
	}
	var items []string
	doc.Find("pre").Each(func(i int, s *goquery.Selection) {
		items = append(items, s.Text())
	})
	return items, nil
}
