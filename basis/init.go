package basis



var(
	Host = Appconf().String("host")
	Port = Appconf().String("port")
	Log = LoggerFile()
)
