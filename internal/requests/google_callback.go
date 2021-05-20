package requests

import (
	"errors"
	"net/http"
)

type CompleteOAuthScheme struct {
	State string
	Code  string
}

func CompleteOAuth(r *http.Request) (CompleteOAuthScheme, error) {
	var request CompleteOAuthScheme

	code := r.FormValue("code")
	if code == "" {
		return request, errors.New("code hasn't been provided")
	}

	state := r.FormValue("code")
	if state == "" {
		return request, errors.New("state hasn't been provided")
	}

	request.Code = code
	request.State = state

	return request, nil
}
