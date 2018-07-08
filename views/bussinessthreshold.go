package views

import(
	"SendAlarm/basis"
	"encoding/json"
	"bytes"
	"net/http"
	"io/ioutil"
	"github.com/pkg/errors"
	"fmt"
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
	Scope int
}

type UpdateResult struct {
	Index string
	Type string
	Id string
	Version string
	Result string
	Shard map[string]int
}

//这里将传入的值进行判断，如果存在就舍弃。
func SearchBussinessItem(b *BussinessThreshold) (total int64,scope int64,err error,idname string){
	var s []term
	s = append(s,term{"term":{"FirstBussiness.keyword":{"value":b.FirstBussiness}}})
	s = append(s,term{"term":{"SecondBussiness.keyword":{"value":b.SecondBussiness}}})
	s = append(s,term{"term":{"ThirdBussiness.keyword":{"value":b.ThirdBussiness}}})
	t := query{"query":{"bool":{"must":s}}}
	bytesDate,err := json.Marshal(t)
	if err != nil {
		basis.Log.Error(err.Error())
		return 0,0,err,""
	}
	reader := bytes.NewReader([]byte(bytesDate))
	request, err := http.NewRequest("POST", basis.Business_url, reader)
	if err != nil {
		basis.Log.Error(err.Error())
		return 0,0,err,""
	}
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		basis.Log.Error(err.Error())
		return 0,0,err,""
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		basis.Log.Error(err.Error())
		return 0,0,err,""
	}
	//str := (*string)(unsafe.Pointer(&respBytes))
	//fmt.Println(*str)
	fmt.Println("Begin ParserTotalJson")
	total,scope,err = ParserTotalJson(respBytes)
	fmt.Println("After ParserTotalJson")
	if err != nil {
		basis.Log.Error(err.Error())
		return 0,scope,err,""
	}
	fmt.Println("Begin ParserIdJson")
	idname,err = ParserIdJson(respBytes)
	fmt.Println("After ParserIdJson")
	if err != nil{
		basis.Log.Error(err.Error())
		return total,scope,err,""
	}
	return total,scope,nil,idname
}

//这里将配置更新到日志中，通过logstash更新到es中
func InsertBussinessItemElasticsearch(b *BussinessThreshold)(err error){
	length,_,err,_ :=SearchBussinessItem(b)
	if length >=1{
		basis.Log.Warning("已存在，跳过")
		return errors.New("已经存在，将忽略")
	}
	if b.Scope ==0{
		DefaultScope,err := basis.Appconf().Int("DefaultScope")
		if err != nil{
			b.Scope = 0
		}
		b.Scope = DefaultScope
	}
	data,err := json.Marshal(b)
	if err != nil{
		basis.Log.Error(err.Error())
		return nil
	}
	basis.Writefile(data,"conf")
	return nil
}

//这里更新es上已经存在的记录，如果不存在，就插入到es中
func UpdateBussinessItemElasticsearch(b *BussinessThreshold)(err error){
	total,_,err,_ := SearchBussinessItem(b)
	if err != err{
		return err
	}
	if total >2 {
		return errors.New("同个业务发现有多个配置")
	}
	if total == 0{
		err := InsertBussinessItemElasticsearch(b)
		if err != nil{
			return err
		}
		return nil
	}
	temp := make(map[string]map[string]int)
	temp["doc"]["Threshold"] = b.Threshold
	if b.Scope !=0 {
		temp["doc"]["Scope"] = b.Scope
	}
	bytesDate,err := json.Marshal(temp)
	if err != nil {
		basis.Log.Error(err.Error())
		return err
	}
	reader := bytes.NewReader([]byte(bytesDate))
	request, err := http.NewRequest("POST", basis.Business_url, reader)
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

//这里得到总数，如果是搜索业务的范围时间，会返回scope进行判断
func ParserTotalJson(Total []byte)(total int64,scope int64,err error){
	var code Code
	err = json.Unmarshal(Total,&code)

	if err != nil {
		return 0,scope,nil
	}
	switch v := code.Hits["total"].(type){
	case int:
		total = int64(v)
	case float64:
		total = int64(float64(v))
		if total == 1{
			switch w := code.Hits["hits"].(type) {
			case []interface{}:
				switch w1 := w[0].(type) {
				case map[string]interface{}:
					switch w2 := w1["_source"].(type) {
					case map[string]interface{}:
						switch w3 := w2["Scope"].(type) {
						case int:
							scope = int64(w3)
							return total,scope,nil
						}
					}
				}	
			default:
				scope = 0
			}
		}
	default:
		total = 0
	}
	return total,scope,err
}

//这里得到Update需要的doc id
func ParserIdJson(body []byte)(id string,err error){
	var code Code
	err = json.Unmarshal(body,&code)
	if err != nil {
		return "",nil
	}
	if code.Hits == nil{
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
			return "",errors.New("hits中数组字段无法解析")
		}

	default:
		return "",errors.New("hits字段无法解析")
	}
	return "",nil
}

//更新后，判断更新是否成功。
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
