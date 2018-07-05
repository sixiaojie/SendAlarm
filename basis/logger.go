package basis
import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/config"
	"encoding/json"
)


//这里需要添加参数值：logfile,maxdays,daily,logAsync

func LoggerFile() (logger *logs.BeeLogger){
	iniconf, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	//这里初始化一个日志配置信息
	log_conf := make(map[string]interface{})
	logfile := iniconf.String("error")
	if logfile == ""{
		logfile = "error.log"
	}
	l := logs.NewLogger()
	//这里将配置信息写入logger中
	log_conf["filename"] = logfile
	maxdays,err := iniconf.Int64("maxdays")
	if err != nil{
		maxdays = 7
	}
	daily := iniconf.String("daily")
	if daily == ""{
		daily = "daily"
	}
	log_conf["daliy"] = daily
	log_conf["maxdays"] = maxdays
	//这里将log的配置转换为字符串类型
	xx,err:= json.Marshal(log_conf)
	if err != nil{
		panic(err)
	}
	l.SetLogger(logs.AdapterFile,string(xx))
	//这里设置日志异步的方式。可以累积写入
	logAsync,err := iniconf.Int64("logAsync")
	l.Async(logAsync)
	return l
}