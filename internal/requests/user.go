package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type UpdateUser struct {
	PicURL string    `json:"picture_url,omitempty"`
	DOB    time.Time `json:"dob,omitempty"`
}

func NewUpdateUser(r *http.Request) (*UpdateUser, error) {
	var updateUser UpdateUser

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &updateUser)
	//TODO: add validation

	return &updateUser, err
}
