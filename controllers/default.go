package controllers

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
	"sort"
)

type MainController struct {
	beego.Controller
}

type WeixinController struct {
	beego.Controller
}

const (
	TOKEN    = "kongtingcom"
	Text     = "text"
	Location = "location"
	Image    = "image"
	Link     = "link"
	Event    = "event"
	Music    = "music"
	News     = "news"
)

func (this *MainController) Get() {
	this.Data["Website"] = "kongting.com"
	this.Data["Email"] = "kongting@kongting.com"
	this.TplNames = "index.tpl"
}

func (this *WeixinController) Get() {
	signature := this.Input().Get("signature")
	beego.Info(signature)
	timestamp := this.Input().Get("timestamp")
	beego.Info(timestamp)
	nonce := this.Input().Get("nonce")
	beego.Info(nonce)
	echostr := this.Input().Get("echostr")
	beego.Info(echostr)
	beego.Info(Signature(timestamp, nonce))
	if Signature(timestamp, nonce) == signature {
		this.Ctx.WriteString(echostr)
	} else {
		this.Ctx.WriteString("")
	}
}

func Signature(timestamp, nonce string) string {
	strs := sort.StringSlice{TOKEN, timestamp, nonce}
	sort.Strings(strs)
	str := ""
	for _, s := range strs {
		str += s
	}
	h := sha1.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}
