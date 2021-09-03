package tags

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLinks(t *testing.T) {
	tests := map[string]func(*testing.T){
		"should return list of valid links": func(t *testing.T) {
			// given
			givenData := `<a href="http://rexegg.com/best-tricks">Greatest Regex Tricks</a>`

			// when
			gotLinks := ParseLinks(givenData)

			// then
			url, err := url.Parse("http://rexegg.com/best-tricks")
			wantLinks := []link{
				{"Greatest Regex Tricks", *url},
			}
			assert.NoError(t, err)
			assert.NotEmpty(t, gotLinks)
			assert.Equal(t, wantLinks, gotLinks)
		},
		"should return empty when no valid links are found": func(t *testing.T) {
			// given
			givenData := `<a href="item?id123123123">Item</a>`
			givenData += `<a href="security.html">Sec</a>`
			givenData += `<a href="mailto:hn@ycombinator.com">Mail</a>`
			givenData += `<a href="ftp://directory.com/test/">FTP Directory</a>`

			// when
			gotLinks := ParseLinks(givenData)

			// then
			assert.Empty(t, gotLinks)
		},
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			run(t)
		})
	}
}
