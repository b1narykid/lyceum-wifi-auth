package requests

import (
	"strings"
	"net/url"
	"net/http"
)

const (
	LoginURL = "https://wlc38.hse.ru/login.html"
)

func NewLogin(user, pass string) (*http.Request, error) {
	data := url.Values{
		"Submit": { "Submit" },
		"buttonClicked": { "4" },
		"username": { user },
		"password": { pass },
	}
	req, err := http.NewRequest(
		"POST", LoginURL,
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}

