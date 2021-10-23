package tags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTitle(t *testing.T) {
	tests := []struct {
		name string
		html string
		want string
	}{
		{
			name: "should return page title when html has valid title tag",
			html: "<html><title>some title</title></html>",
			want: "some title",
		},
		{
			name: "should empty title when html has invalid title tag",
			html: "<html><title>some name</section></html>",
			want: "",
		},
		{
			name: "should empty title when html does not have title tag",
			html: "<html><section>some section</section></html>",
			want: "",
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got := ParseTitle(testCase.html)
			assert.Equal(t, testCase.want, got)
		})
	}
}
