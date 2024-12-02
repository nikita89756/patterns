package handler

import "encoding/json"

type JSONHandler struct{
	nextHandler *ParseHandler
}

func (j *JSONHandler)Handle(s string)(string){
	var js json.RawMessage
	if json.Unmarshal([]byte(s),&js) == nil{
		return "json"
	}
	handler:=*j.nextHandler
	return handler.Handle(s)
}

func (j *JSONHandler) SetNext(h ParseHandler){
	j.nextHandler = &h
}