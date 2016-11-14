package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"survey_api/app/db"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) RegisterAccount() revel.Result {
	//create temporary user variable
	//set response type
	var user db.User
	c.Response.ContentType = "application/json; charset=UTF-8"

	//read body and defer close
	body, err := ioutil.ReadAll(io.LimitReader(c.Request.Body, 1048576))
	defer c.Request.Body.Close()

	//handle error
	if err != nil {
		panic(err)
	}

	//unmarshal body
	if err := json.Unmarshal(body, &user); err != nil {
		c.Response.Status = 422
	}
	// register and save account
	user.RegisterAccount()
	c.Response.Status = http.StatusCreated
	return c.RenderJson(&user)
}
func (c App) Login() revel.Result {
	var user db.User
	c.Response.ContentType = "application/json; charset=UTF-8"

	//read body and defer close
	body, err := ioutil.ReadAll(io.LimitReader(c.Request.Body, 1048576))
	defer c.Request.Body.Close()

	//handle error
	if err != nil {
		panic(err)
	}

	//unmarshal body
	if err := json.Unmarshal(body, &user); err != nil {
		c.Response.Status = 422
	}
	return c.RenderJson(user.Login())
}
