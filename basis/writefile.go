package basis

import (
	"time"
	"os"
)

func Writefile(msg []byte,item string){
	config := Appconf()
	Send := config.String(FindLog(item))
	if Send == ""{
		Send = FindLog(item)+".log"
	}
	fd,_:=os.OpenFile(Send,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	//buf:=[]byte(msg)
	fd.Write(msg)
	fd.Write([]byte("\n"))
	fd.Close()
}


func FindLog(item string) string{
	if item == "conf"{
		return "email_json"
	}else{
		return "access"
	}
}


//这里得到业务日志的索引名字
func GetIndexName() string{
	now := time.Now().Format("2006.01")
	return Appconf().String("AlarmConfIndexName")+"-"+now
}


func GetBussinessConfIndexName() string{
	return Appconf().String("BussinessConfIndexName")
}
