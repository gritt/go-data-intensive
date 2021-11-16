package http

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetContents(t *testing.T) {
	ctx := context.TODO()

	givenURL := url.URL{
		Scheme: "http",
		Host:   "localhost",
		Path:   "mock",
	}

	tests := map[string]func(*testing.T, *mockHttpClient){
		"should return error when fails to get": func(t *testing.T, mockHttpClient *mockHttpClient) {
			// given
			givenErr := errors.New("get error")
			mockHttpClient.On("Get", givenURL.String()).Return(new(http.Response), givenErr)

			browser := NewBrowser(mockHttpClient)

			// when
			contents, err := browser.GetContents(ctx, givenURL)

			// then
			assert.Empty(t, contents)
			assert.EqualError(t, err, "request error: get error")
		},
		"should return error when invalid status code": func(t *testing.T, mockHttpClient *mockHttpClient) {
			// given
			givenRespBody := new(mockHttpResponse)
			givenRespBody.On("Read", mock.Anything).Return(0, nil)
			givenRespBody.On("Close").Return(nil)

			givenResp := http.Response{
				StatusCode:    403,
				Body:          givenRespBody,
				ContentLength: 0,
			}

			mockHttpClient.On("Get", givenURL.String()).Return(&givenResp, nil)

			browser := NewBrowser(mockHttpClient)

			// when
			contents, err := browser.GetContents(ctx, givenURL)

			// then
			assert.Empty(t, contents)
			assert.EqualError(t, err, "invalid status code: 403")
		},
		"should return error when page is empty": func(t *testing.T, mockHttpClient *mockHttpClient) {
			// given
			givenRespBody := new(mockHttpResponse)
			givenRespBody.On("Read", mock.Anything).Return(0, nil)
			givenRespBody.On("Close").Return(nil)

			givenResp := http.Response{
				StatusCode:    200,
				Body:          givenRespBody,
				ContentLength: 0,
			}

			mockHttpClient.On("Get", givenURL.String()).Return(&givenResp, nil)

			browser := NewBrowser(mockHttpClient)

			// when
			contents, err := browser.GetContents(ctx, givenURL)

			// then
			assert.Empty(t, contents)
			assert.EqualError(t, err, "empty body: 0")
		},
		"should return error when failed to read contents": func(t *testing.T, mockHttpClient *mockHttpClient) {
			// given
			givenErr := errors.New("response error")
			givenRespBody := new(mockHttpResponse)
			givenRespBody.On("Read", mock.Anything).Return(0, givenErr)
			givenRespBody.On("Close").Return(nil)

			givenResp := http.Response{
				StatusCode:    200,
				Body:          givenRespBody,
				ContentLength: 5,
			}

			mockHttpClient.On("Get", givenURL.String()).Return(&givenResp, nil)

			browser := NewBrowser(mockHttpClient)

			// when
			contents, err := browser.GetContents(ctx, givenURL)

			// then
			assert.Empty(t, contents)
			assert.EqualError(t, err, "read error: response error")
		},
		"should return contents with success": func(t *testing.T, mockHttpClient *mockHttpClient) {
			// given
			givenRespContent := "hello"
			givenRespBody := bytes.NewReader([]byte(givenRespContent))
			givenResp := http.Response{
				StatusCode:    200,
				Body:          ioutil.NopCloser(givenRespBody),
				ContentLength: 5,
			}

			mockHttpClient.On("Get", givenURL.String()).Return(&givenResp, nil)

			browser := NewBrowser(mockHttpClient)

			// when
			contents, err := browser.GetContents(ctx, givenURL)

			// then
			assert.NoError(t, err)
			assert.Equal(t, givenRespContent, contents)
		},
	}

	for name, run := range tests {
		t.Run(name, func(t *testing.T) {
			mockHttpClient := new(mockHttpClient)

			run(t, mockHttpClient)
		})
	}
}

type mockHttpClient struct {
	mock.Mock
}

func (m *mockHttpClient) Get(url string) (*http.Response, error) {
	args := m.Called(url)
	return args.Get(0).(*http.Response), args.Error(1)
}

type mockHttpResponse struct {
	mock.Mock
}

func (m *mockHttpResponse) Read(p []byte) (n int, err error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)

}

func (m *mockHttpResponse) Close() error {
	args := m.Called()
	return args.Error(0)
}
