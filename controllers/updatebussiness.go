package controllers
import (
	"github.com/astaxie/beego"
	"SendAlarm/views"
	"encoding/json"
)

type BuissnessController struct {
	beego.Controller
}

func (u *BuissnessController) Insert(){
	var buss views.BussinessThreshold
	var sr views.Base
	SR := make(map[string]interface{})
	json.Unmarshal(u.Ctx.Input.RequestBody, &buss)
	err := views.InsertBussinessItemElasticsearch(&buss)
	if err != nil {
		SR["message"] = err.Error()
		sr = views.SendReturn(sr, 4022, "Insert failed", SR)
	} else{
		SR["message"] = "success"
		sr = views.SendReturn(sr,0,"ok",SR)
	}
	u.Data["json"] = sr
	u.ServeJSON()
}


func (u *BuissnessController) Update(){
	var buss views.BussinessThreshold
	var sr views.Base
	SR := make(map[string]interface{})
	err := views.UpdateBussinessItemElasticsearch(&buss)
	if err != nil {
		SR["message"] = err.Error()
		sr = views.SendReturn(sr, 4023, "Update failed", SR)
	} else{
		SR["message"] = "success"
		sr = views.SendReturn(sr,0,"ok",SR)
	}
	u.Data["json"] = sr
	u.ServeJSON()
}

