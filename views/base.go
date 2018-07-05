package views

type Base struct {
	Version string
	Code int
	Status string
	Msg map[string]interface{}
}

func SendReturn(b Base,code int,status string,msg map[string]interface{})(Base){
	b.MakeReturnVersion()
	b.Code = code
	b.Status = status
	b.Msg = msg
	return b
}

func (b *Base)MakeReturnVersion(){
	version := "xxxx"
	b.Version = version
}
