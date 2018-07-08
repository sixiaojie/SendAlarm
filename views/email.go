package views

import (
	"fmt"
	"github.com/pkg/errors"
	"SendAlarm/basis"
	"encoding/json"
	"bytes"
	"net/http"
	"io/ioutil"
	"strconv"
	"time"
)

type EmailType struct {
	Id int
	User []string
	Cc []string
	Alias string
	Host string
	Status string
	Msg string
	Subject string
	Message string
	Localtime int64
	FirstBussiness string
	SecondBussiness string
	ThirdBussiness string
	Extra map[string]string
}

func (e *EmailType) DealExtra(){
	if len(e.Extra) == 0{
		fmt.Println("xxxx")
	}
}
//这里将发邮件和存储的放在一起
func SendInit(e EmailType) (errcode int,err error){
	err = CheckJsonData(&e)
	if err != nil {
		return -1,err
	}
	_,err = SendEmails(&e)
	if err != nil{
		return 10001,err
	}
	err = StoreElasticsearch(&e)
	if err != nil {
		return 10002,err
	}
	return 0,nil
}
// 这里与阈值对比，然后决定发邮件
func SendEmails(e *EmailType)(code int64,err error){
	thres := &BussinessThreshold{e.FirstBussiness,e.SecondBussiness,e.ThirdBussiness,0,0}
	_,scope,threshold,err,_ := SearchBussinessItem(thres)
	if err != nil{
		return 6044,err
	}
	total,err := SearchElasticsearch(e,scope)
	if err != nil{
		return total,err
	}
	Email := basis.Email_Server{}
	if total ==0{
		Email.SendMail(basis.Alarm,e.Subject,e.Message,e.User,e.Cc,basis.Logger)
		return 0,err
	}else if  total>=threshold{
		Email.SendMail(basis.Alarm,e.Subject,e.Message,e.User,e.Cc,basis.Logger)
		return 0,err
	}
	return 0,err
}

//这里存储信息到Elasticsearch
func StoreElasticsearch(e *EmailType)(err error){
	err = e.analyse()
	if err != nil{
		return err
	}
	e.Localtime = time.Now().Unix()
	data,err := json.Marshal(e)
	if err != nil{
		basis.Log.Error(err.Error())
		return err
	}
	basis.Writefile(data,"access")
	return nil
}

//这里判断接受的json是否正规
func CheckJsonData(e *EmailType)(err error){
	return err
}

//这里获得出现的次数，根据业务，业务的设置时间，阈值。0代表发邮件，0代表不发邮件
func SearchElasticsearch(e *EmailType,scope int64)(times int64,err error){
	var s []term
	err = e.analyse()
	if err != nil{
		return 0,err
	}
	//这里将时间加上
	localtime := time.Now().Unix() - scope*60
	s = append(s,term{"term":{"FirstBussiness.keyword":{"value":e.FirstBussiness}}})
	s = append(s,term{"term":{"SecondBussiness.keyword":{"value":e.SecondBussiness}}})
	s = append(s,term{"term":{"ThirdBussiness.keyword":{"value":e.ThirdBussiness}}})
	s = append(s,term{"term":{"ThirdBussiness.keyword":{"value":e.ThirdBussiness}}})
	s = append(s,term{"term":{"Status.keyword":{"value":e.Status}}})
	s = append(s,term{"term":{"Msg.keyword":{"value":e.Msg}}})
	s = append(s,term{"range":{"Localtime":{"gt":strconv.FormatInt(localtime,10)}}})
	t := query{"query":{"bool":{"must":s}}}
	bytesDate,err := json.Marshal(t)
	if err != nil {
		basis.Log.Error(err.Error())
		return 5047,err
	}
	reader := bytes.NewReader([]byte(bytesDate))
	request, err := http.NewRequest("POST", basis.Alarm_url, reader)
	if err != nil {
		basis.Log.Error(err.Error())
		return 5046,err
	}
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		basis.Log.Error(err.Error())
		return 5045,err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		basis.Log.Error(err.Error())
		return 5044,err
	}
	times,scope,_,err = ParserTotalJson(respBytes)
	if err != nil{
		return 5042,err
	}
	return times,nil
}

func (e *EmailType) analyse()(error){
	status,host,msg,length := Decompose(e.Subject)
	if length != 3{
		return errors.New("Subject格式不正确")
	}
	e.Host = host
	e.Msg = msg
	e.Status = status
	return nil
}
