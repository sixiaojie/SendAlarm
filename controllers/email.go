package controllers
import (
"github.com/astaxie/beego"
	"encoding/json"
	"SendAlarm/views"
	"fmt"
)

type EmailController struct {
	beego.Controller
}
//这里接受信息，并且返回信息。
func (e *EmailController) Send(){
	var eml views.EmailType
	var sr views.Base
	SR := make(map[string]interface{})
	json.Unmarshal(e.Ctx.Input.RequestBody, &eml)
	fmt.Println("before sendmail")
	code,err := views.SendInit(eml)
	fmt.Println("after sendmail")
	if err != nil{
		SR["message"] = err.Error()
		sr = views.SendReturn(sr,code,"Failed",SR)
	}else{
		SR["message"] = "success"
		sr = views.SendReturn(sr,0,"ok",SR)
	}
	e.Data["json"] = sr
	e.ServeJSON()
}
