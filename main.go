package main

import (
	_ "email/routers"

	"github.com/astaxie/beego"
	"fmt"
)

func main() {
	fmt.Println(beego.BConfig)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
