package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewTask(r *http.Request) (*Task, error) {
	var task Task

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &task)
	//TODO: add validation

	return &task, err
}
