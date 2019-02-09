package controllers

import (
	"encoding/json"
	"fmt"
	"serveur/models"

	"github.com/astaxie/beego"
)

type Registration struct {
	beego.Controller
}

//
func (c *Registration) Get() {
	c.TplName = "desktop/registration.html"
}

func (c *Registration) Post() {

}

// Post EmailRegistration handle post request for new user
// email registration
func (c *AuthController) Post() {
	var user models.UserData
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)

	err = user.Authenticate()
	if err != nil {
		fmt.Println(err)
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = struct {
			Error string `json:"error"`
		}{err.Error()}
		c.ServeJSON()
		return
	}

	err = user.CreateSession(c.Ctx.Input.Context)
	if err != nil {
		fmt.Println(err)
		c.Ctx.Output.SetStatus(500)
		return
	}
	c.Data["json"] = user
	c.ServeJSON()
}

// Post method of EmailLoginController
// expect user json data in request body
// authenticate user and create new session
// for user
// func (c *EmailLoginController) Post() {
// 	var user models.User
// 	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
// 	if err != nil {
// 		fmt.Println(err)
// 		c.Ctx.Output.SetStatus(401)
// 		return
// 	}

// 	err = user.Authenticate()
// 	if err != nil {
// 		fmt.Println(err)
// 		c.Ctx.Output.SetStatus(401)
// 		return
// 	}

// 	err = user.CreateSession(c.Ctx.Input.Context)
// 	if err != nil {
// 		fmt.Println(err)
// 		c.Ctx.Output.SetStatus(500)
// 		return
// 	}
// 	c.Data["json"] = user
// 	c.ServeJSON()
// }

// Delete handler for when user send request to logout
func (c *LogoutController) Delete() {
	models.DestroySession(c.Ctx.Input.Context)
	c.Finish()
}
