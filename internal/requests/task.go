package requests

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type NewTask struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetTask struct {
	ID int `uri:"id" binding:"required"`
}

func CreateTask(r *http.Request) (*NewTask, error) {
	var task NewTask

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &task)
	//TODO: add validation

	return &task, err
}

func GetOne(c *gin.Context) (*GetTask, error) {
	var task GetTask
	err := c.BindUri(&task)
	return &task, err
}
