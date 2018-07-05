package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["email/controllers:ObjectController"] = append(beego.GlobalControllerRouter["email/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["email/controllers:ObjectController"] = append(beego.GlobalControllerRouter["email/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["email/controllers:ObjectController"] = append(beego.GlobalControllerRouter["email/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["email/controllers:ObjectController"] = append(beego.GlobalControllerRouter["email/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["email/controllers:ObjectController"] = append(beego.GlobalControllerRouter["email/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["email/controllers:UserController"] = append(beego.GlobalControllerRouter["email/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["email/controllers:UserController"] = append(beego.GlobalControllerRouter["email/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["email/controllers:UserController"] = append(beego.GlobalControllerRouter["email/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["email/controllers:UserController"] = append(beego.GlobalControllerRouter["email/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["email/controllers:UserController"] = append(beego.GlobalControllerRouter["email/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["email/controllers:UserController"] = append(beego.GlobalControllerRouter["email/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["email/controllers:UserController"] = append(beego.GlobalControllerRouter["email/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["email/controllers:NameController"] = append(beego.GlobalControllerRouter["email/controllers:NameController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/name`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SendAlarm/controllers:BuissnessController"] = append(beego.GlobalControllerRouter["SendAlarm/controllers:BuissnessController"],
		beego.ControllerComments{
			Method: "Insert",
			Router: `/insert`,
			AllowHTTPMethods: []string{"POST"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["SendAlarm/controllers:BuissnessController"] = append(beego.GlobalControllerRouter["SendAlarm/controllers:BuissnessController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/update`,
			AllowHTTPMethods: []string{"POST"},
			MethodParams: param.Make(),
			Params: nil})

}
