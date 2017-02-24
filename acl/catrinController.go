package acl

import (
	"github.com/astaxie/beego"
)

type CatrinController struct {
	beego.Controller

	Auth
}

func (this *CatrinController) Prepare() {

}
func Decode(row string, ctlr *CatrinController) error {

	return nil
}

func Encode(ctlr *CatrinController) string {

	return ""
}
