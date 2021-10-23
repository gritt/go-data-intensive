package tags

import (
	"errors"
	"net/url"
	"regexp"
	"strings"
)

type link struct {
	name string
	url  url.URL
}

func ParseLinks(html string) (links []link) {
	tags := hyperlinkTags(html)
	if len(tags) == 0 {
		return
	}

	for _, hyperlink := range tags {
		hrefURL, err := hyperlinkToURL(hyperlink)
		if err != nil {
			continue
		}
		if !isValidURL(hrefURL) {
			continue
		}
		name, err := hyperlinkToName(hyperlink)
		if err != nil {
			continue
		}

		links = append(links, link{name: name, url: hrefURL})
	}

	return
}

func hyperlinkTags(html string) []string {
	parser, err := regexp.Compile(`(<a .*?href=.*?"(.*?)"(.|\n)*?>((.|\n)*?)<.*?/a.*?>)`)
	if err != nil {
		// TODO log err
		return []string{}
	}
	return parser.FindAllString(html, len(html))
}

func hyperlinkToURL(hrefTag string) (url.URL, error) {
	parser, err := regexp.Compile(`(href\s*=\s*(?:"|')(.*?)(?:"|'))`)
	if err != nil {
		// TODO log err
		return url.URL{}, err
	}

	hrefAttribute := parser.Find([]byte(hrefTag))
	if len(hrefAttribute) == 0 {
		// TODO log err
		return url.URL{}, err
	}

	hrefValue := strings.Replace(string(hrefAttribute), `href=`, ``, 1)
	hrefValue = strings.Replace(hrefValue, `"`, ``, 2)
	hrefValue = strings.Replace(hrefValue, `'`, ``, 2)

	hrefURL, err := url.Parse(hrefValue)
	if err != nil {
		// TODO log err
		return url.URL{}, err
	}

	return *hrefURL, nil
}

func hyperlinkToName(hrefTag string) (string, error) {
	parser, err := regexp.Compile(`<a.*?>\s*(.*?)\s*</a>`)
	if err != nil {
		// TODO log err
		return "", err
	}

	data := parser.FindSubmatch([]byte(hrefTag))
	if len(data) < 2 {
		// TODO log err
		return "", errors.New("some err")
	}
	name := strings.TrimSpace(string(data[1]))

	return name, nil
}

func isValidURL(hrefURL url.URL) bool {
	if !hrefURL.IsAbs() {
		return false
	}
	if hrefURL.Scheme != "http" && hrefURL.Scheme != "https" {
		return false
	}
	return true
}
