package views

import "fmt"

type EmailType struct {
	Id int
	User []string
	Cc []string
	Alias string
	Subject string
	Message string
	Bussiness string
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
	err = SendEmails(&e)
	if err != nil{
		return 10001,err
	}
	err = StoreElasticsearch(&e)
	if err != nil {
		return 10002,err
	}
	return 0,nil
}
// 这里发邮件
func SendEmails(e *EmailType)(err error){
	return nil
}

//这里存储信息到Elasticsearch
func StoreElasticsearch(e *EmailType)(err error){
	return nil
}

//这里判断接受的json是否正规
func CheckJsonData(e *EmailType)(err error){
	return err
}

//这里获得出现的次数，根据业务，业务的设置时间，阈值。0代表发邮件，0代表不发邮件
func SearchElasticsearch(e *EmailType)(times int,err error){
	return 0,err
}