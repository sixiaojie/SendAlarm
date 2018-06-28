package views

type Base struct {
	Version string `json:"version"`
	Code int `json:"code"`
	Status string `json:"status"`
	Msg map[string]interface{} `json:"msg"`
}

func SendReturn(b Base,code int,status string,msg map[string]interface{}){
	b.MakeReturnVersion()
	b.Code = code
	b.Status = status
	b.Msg = msg
}

func (b *Base)MakeReturnVersion(){
	version := "xxxx"
	b.Version = version
}
