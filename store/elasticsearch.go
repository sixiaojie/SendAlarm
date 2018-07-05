package store

import (
	"reflect"
	"SendAlarm/views"
	"encoding/json"
	"fmt"
)

type elatico struct {
	Id int
	User map[string][]string
	Cc map[string][]string
	Alias string
	Status string
	Host string
	Problem string
	Message string
	Bussiness string
	Extra map[string]string
}

func (l *elatico) Init(e *views.EmailType)(length int){
	status,host,message,length := views.Decompose(e.Subject)
	l.Status = status
	l.Host = host
	l.Problem = message
	return length
}

//这里将需要发送的日志写入到日志（这个日志存储邮件发送的基本的信息）。
func InsertElatico(l *elatico,e  *views.EmailType){
	l.Init(e)
	temp := make(map[string]interface{})
	s := reflect.ValueOf(l).Elem()
	typeOFL := s.Type()
	for i :=0;i<s.NumField();i++{
		temp[typeOFL.Field(i).Name] = s.Field(i).Interface()
	}
	msg,err := json.Marshal(temp)
	if err != nil{
		//这里写入程序日志中
	}
	//这里写入到业务日志中
	fmt.Println(msg)
}

//这里将
func SearchElatico(l *elatico,e  *views.EmailType)(int64){
	length := l.Init(e)
	if length !=3{
		return 9999
	}

}