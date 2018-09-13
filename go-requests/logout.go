package requests

import (
	"strings"
	"net/url"
	"net/http"
)

const (
	LogoutURL = "https://wlc38.hse.ru/logout.html"
)

func NewLogout() (*http.Request, error) {
	data := url.Values{
		"Logout": {
			"Logout",
		},
	}
	req, err := http.NewRequest(
		"POST", LogoutURL,
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}
