package tags

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLinks(t *testing.T) {
	tests := map[string]func(*testing.T){
		"should return list of links": func(t *testing.T) {
			// given
			givenData := `<a href="http://rexegg.com/regex-best-trick.html">The Greatest Regex Trick Ever</a>`
			givenData += `<a href="http://website.com">Gilv.es</a>`

			givenData += `<a href="item?id123123123">Item</a>`
			givenData += `<a href="security.html">Sec</a>`

			givenData += `<a href="mailto:hn@ycombinator.com">Mail</a>`

			givenData += `<a href="ftp://directory.com/test/">FTP Directory</a>`

			// when
			gotLinks := ParseLinks(givenData)

			// then
			url, err := url.Parse("http://rexegg.com/regex-best-trick.html")
			wantLinks := []link{
				{"The Greatest Regex Trick Ever (2014)", *url},
			}
			assert.NoError(t, err)
			assert.NotEmpty(t, gotLinks)
			assert.Equal(t, wantLinks, gotLinks)
		},
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}
