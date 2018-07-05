package controllers
import (
	"github.com/astaxie/beego"
	"encoding/json"
	"SendAlarm/views"
)

type ThresholdController struct {
	beego.Controller
}
//这里接受信息，并且返回信息。
func (e *ThresholdController) Post(){
	var eml views.EmailType
	var sr views.Base
	SR := make(map[string]interface{})
	json.Unmarshal(e.Ctx.Input.RequestBody, &eml)
	code,err := views.SendInit(eml)
	if err != nil{
		SR["message"] = err.Error()
		views.SendReturn(sr,code,"Failed",SR)
	}else{
		SR["message"] = "success"
		views.SendReturn(sr,0,"ok",SR)
	}
	e.Data["json"] = sr
}

