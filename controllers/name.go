package controllers
import (
	"github.com/astaxie/beego"
)

type NameController struct {
	beego.Controller
}

func (u *NameController) Get(){
	temp := make(map[string]string)
	temp["name"] = "sijie"
	temp["sex"] = "man"
	u.Data["json"] = temp
	u.ServeJSON()
}