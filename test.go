package main

import (
	"SendAlarm/views"
	"fmt"
)

func main(){
	b := &views.BussinessThreshold{"fk","backend","api",10,10}
	total,scope,threshold,err,idname := views.SearchBussinessItem(b)
	fmt.Println(total,scope,threshold,err,idname)
}
