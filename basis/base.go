package basis

import "github.com/astaxie/beego/config"

func Appconf()(config.Configer){
	iniconf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil{
		panic(err)
	}
	return iniconf
}