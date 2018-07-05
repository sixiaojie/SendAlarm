package views

import(
	"SendAlarm/basis"
	"encoding/json"
	"bytes"
	"net/http"
	"io/ioutil"
	"github.com/pkg/errors"
)

type value map[string]string
type field map[string]value
type term map[string]field

type must map[string][]term
type Bool1 map[string]must
type query map[string]Bool1


type shard map[string]int
type hit map[string]interface{}
type Code struct {
	Took int `json:took`
	TimeOut bool `json:time_out`
	_Shards shard `json:shards`
	Hits   hit 	`json:hits`
}

type BussinessThreshold struct {
	FirstBussiness string
	SecondBussiness string
	ThirdBussiness string
	Threshold int
}

type UpdateResult struct {
	Index string
	Type string
	Id string
	Version string
	Result string
	Shard map[string]int
}


func SearchBussinessItem(b *BussinessThreshold) (total int64,err error,idname string){
	indexname := basis.GetBussinessConfIndexName()
	url := "http://"+basis.Host+":"+basis.Port+"/"+indexname+"/_search"
	var s []term
	s = append(s,term{"term":{"FirstBussiness.keyword":{"value":b.FirstBussiness}}})
	s = append(s,term{"term":{"SecondBussiness.keyword":{"value":b.SecondBussiness}}})
	s = append(s,term{"term":{"ThirdBussiness.keyword":{"value":b.ThirdBussiness}}})
	t := query{"query":{"bool":{"must":s}}}
	bytesDate,err := json.Marshal(t)
	if err != nil {
		basis.Log.Error(err.Error())
		return 0,err,""
	}
	reader := bytes.NewReader([]byte(bytesDate))
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		basis.Log.Error(err.Error())
		return 0,err,""
	}
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		basis.Log.Error(err.Error())
		return 0,err,""
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		basis.Log.Error(err.Error())
		return 0,err,""
	}
	//str := (*string)(unsafe.Pointer(&respBytes))
	//fmt.Println(*str)
	total,err = ParserTotalJson(respBytes)
	if err != nil {
		basis.Log.Error(err.Error())
		return 0,err,""
	}
	idname,err = ParserIdJson(respBytes)
	if err != nil{
		basis.Log.Error(err.Error())
		return total,err,""
	}
	return total,nil,idname
}

//这里将配置更新到日志中，通过logstash更新到es中
func InsertBussinessItemElasticsearch(b *BussinessThreshold)(err error){
	data,err := json.Marshal(b)
	if err != nil{
		basis.Log.Error(err.Error())
	}
	basis.Writefile(data,"email_json")
	return nil
}

//这里更新es上已经存在的记录，如果不存在，就插入到es中
func UpdateBussinessItemElasticsearch(b *BussinessThreshold)(err error){
	total,err,idname := SearchBussinessItem(b)
	if err != err{
		return err
	}
	if total !=1 {
		return errors.New("同个业务发现有多个配置")
	}
	indexname := basis.GetBussinessConfIndexName()
	url := "http://"+basis.Host+":"+basis.Port+"/"+indexname+"/"+idname+"/_update"
	temp := make(map[string]map[string]int)
	temp["doc"]["Threshold"] = b.Threshold
	bytesDate,err := json.Marshal(temp)
	if err != nil {
		basis.Log.Error(err.Error())
		return err
	}
	reader := bytes.NewReader([]byte(bytesDate))
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		basis.Log.Error(err.Error())
		return err
	}
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		basis.Log.Error(err.Error())
		return err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		basis.Log.Error(err.Error())
		return err
	}
	err = UpdateThresholdReturnResult(respBytes)
	if err != nil{
		return err
	}
	return nil
}

//这里得到总数
func ParserTotalJson(Total []byte)(total int64,err error){
	var code Code
	err = json.Unmarshal(Total,&code)
	if err != nil {
		return 0,nil
	}
	switch v := code.Hits["total"].(type){
	case int:
		total = int64(v)
	case float64:
		total = int64(float64(v))
	default:
		total = 0
	}
	return total,err
}

//这里得到Update需要的doc id

func ParserIdJson(body []byte)(id string,err error){
	var code Code
	err = json.Unmarshal(body,&code)
	if err != nil {
		return "",nil
	}
	switch v:= code.Hits["hits"].(type) {
	case []interface {}:
		length := len(v)
		if length != 1{
			return "",errors.New("长度超过限制")
		}
		switch w := v[0].(type){
		case map[string]interface{}:
			switch vv := w["_id"].(type) {
			case string:
				return vv,err
			default:
				return "",errors.New("id字段不是string，无法解析")
			}
		default:
			return "",errors.New("hits字段无法解析")
		}

	default:
		return "",errors.New("hits字段无法解析")
	}
	return "",nil
}


func UpdateThresholdReturnResult(b []byte)(err error){
	var updateresult UpdateResult
	err = json.Unmarshal(b,&updateresult)
	if err != nil{
		return err
	}
	if updateresult.Result == "updated"{
		return nil
	}
	return errors.New("修改失败")
}