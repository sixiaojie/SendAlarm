package views

import "strings"

//将subject分解成的字段
func Decompose(subject string)(status,host,msg string,length int){
	temp := strings.Split(subject,":")
	if len(temp) == 3{
		return temp[0],temp[1],temp[2],len(temp)
	}
	if len(temp) == 2{
		return temp[0],"",temp[1],len(temp)
	}
	return "","","",len(temp)
}
