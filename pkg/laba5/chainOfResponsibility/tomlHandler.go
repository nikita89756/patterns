package handler

import (

	"github.com/BurntSushi/toml"
)

type TomlHandler struct{
	nextHandler *ParseHandler
}

func (j *TomlHandler)Handle(s string)(string){
	var data interface{}
	err := toml.Unmarshal([]byte(s), &data)
	if err == nil {
		return "toml"
	}
	handler:=*j.nextHandler
	return handler.Handle(s)
}

func (y *TomlHandler) SetNext(h ParseHandler){
	y.nextHandler = &h
}