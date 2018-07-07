package basis



var(
	Host = Appconf().String("host")
	Port = Appconf().String("port")
	Alarm = Appconf().String("Alias")
	Log = LoggerFile()
	Business_url = "http://"+Host+":"+Port+"/"+GetBussinessConfIndexName()+"/_search"
	Alarm_url = "http://"+Host+":"+Port+"/"+GetIndexName()+"/_search"
	Logger = LoggerFile()
)
