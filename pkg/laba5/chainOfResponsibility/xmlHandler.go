package handler

import (
	"bytes"
	"encoding/xml"
)

type XmlHandler struct{
	nextHandler *ParseHandler
}

func (h *XmlHandler)Handle(s string)(string){
	var data interface{}
	decoder := xml.NewDecoder(bytes.NewReader([]byte(s)))
	err := decoder.Decode(&data)
	if err == nil {
		return "xml"
	}
	handler:=*h.nextHandler
	return handler.Handle(s)
}

func (p *XmlHandler) SetNext(h ParseHandler){
	p.nextHandler = &h
}

