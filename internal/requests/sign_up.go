package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type SignUp struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewSignUp(r *http.Request) (*SignUp, error) {
	var request SignUp

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		return nil, err
	}

	return &request, nil
}