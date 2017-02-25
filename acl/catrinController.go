package acl

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"

	"github.com/astaxie/beego"

	"fmt"
	"strings"
)

type CatrinController struct {
	beego.Controller

	Auth

	Token string
}

func (this *CatrinController) Prepare() {

	row := this.Ctx.Input.Header("AUTH-X")

	if row == "" {
		this.Abort("403")
		this.StopRun()
	}

	Decode(row, this)

	fmt.Println(this.Ctx.Output.Status)
}

func (this *CatrinController) Finish() {

}

func Decode(row string, ctlr *CatrinController) error {

	auth := strings.Split(row, ".")

	ctlr.Token = beego.AppConfig.String(auth[0])

	if ctlr.Token == "" {
		ctlr.Abort("403")
		ctlr.StopRun()
	}

	return nil
}

func (ctlr *CatrinController) Encode() {

	body, _ := json.Marshal(ctlr.Auth)

	block, err := aes.NewCipher([]byte(ctlr.Token))
	if err != nil {

		ctlr.StopRun()
	}

	ciphertext := make([]byte, aes.BlockSize+len(body))

	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		ctlr.StopRun()
	}

	stream := cipher.NewCFBEncrypter(block, iv)

	stream.XORKeyStream(ciphertext[aes.BlockSize:], body)

	fmt.Println(base64.URLEncoding.EncodeToString(ciphertext))
	ctlr.Ctx.Output.Header("AUTH-X", base64.URLEncoding.EncodeToString(ciphertext))

}
