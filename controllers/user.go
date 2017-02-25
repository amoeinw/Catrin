package controllers

import (
	"Amoein/Catrin/acl"
)

// Operations about Users
type UserController struct {
	acl.CatrinController
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {

	u.Ctx.Output.Body([]byte("ok"))
	u.Encode()
	u.ServeJSON()

}
