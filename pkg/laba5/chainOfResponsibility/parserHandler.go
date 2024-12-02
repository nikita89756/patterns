package handler

type ParseHandler interface{
	SetNext(ParseHandler)
	Handle(string) (string)
}

type BaseParseHandler struct{
	nextHandler *ParseHandler
}

func (p *BaseParseHandler) SetNext(h ParseHandler){
	p.nextHandler = &h
}

func (p *BaseParseHandler) Handle(s string)(string){
	if p.nextHandler==nil{
		return "string"
	}
	handler:=*p.nextHandler
	return handler.Handle(s)
}