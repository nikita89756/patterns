package messagehandler

import (
	"fmt"
	"strings"
)

type MessageHandler interface{
	ProcessMessage(string)
	GetByte()[]byte
	GetString() string
	set(string)
}

type BaseMessageDecorator struct{
	message string
}

func (m *BaseMessageDecorator) ProcessMessage(msg string){
	m.message = msg
}

func (m *BaseMessageDecorator) GetByte()[]byte{
	return []byte(m.message)
}

func (m *BaseMessageDecorator) GetString()string{
	return m.message
}

func (m *BaseMessageDecorator) set(msg string){
	m.message = msg
}



type EmojiMessageDecorator struct{
	Handler MessageHandler
}

func (m *EmojiMessageDecorator) ProcessMessage(msg string){
	var i int
	cnt:=true
	m.Handler.ProcessMessage(msg)
	str := m.Handler.GetString()
	for cnt{
		if i = strings.Index(str,":)");i!=-1{
			str  = str[:i] +"😃"+ str[i+2:]
			cnt = true
			continue
		}

		if i = strings.Index(str,":(");i!=-1{
			str  = str[:i] +"😟"+ str[i+2:]
			cnt = true
			continue
		}

		if i = strings.Index(str,":D");i!=-1{
			str  = str[:i] +"😁"+ str[i+2:]
			cnt = true
			continue
		}

		cnt = false

	}
	m.Handler.set(str)
}

func (m *EmojiMessageDecorator) GetByte()[]byte{
	return m.Handler.GetByte()
}

func (m *EmojiMessageDecorator) GetString()string{
	return m.Handler.GetString()
}

func (m *EmojiMessageDecorator) set(msg string){
	m.Handler.set(msg)
}

type LinkMessageDecorator struct{
	Handler MessageHandler
}

func (m *LinkMessageDecorator) ProcessMessage(msg string){
	m.Handler.ProcessMessage(msg)
	str := m.Handler.GetString()
	if strings.Index(str,"http://") !=-1|| strings.Index(str,"https://") !=-1{
		str = fmt.Sprintf("|Данное сообщение может содержать вредоносную ссылку|\n" + str)
	}
	m.Handler.set(str)
}

func (m *LinkMessageDecorator) GetByte()[]byte{
	return m.Handler.GetByte()
}

func (m *LinkMessageDecorator) GetString()string{
	return m.Handler.GetString()
}

func (m *LinkMessageDecorator) set(msg string){
	m.set(msg)
}

func CreateMessageHandler(level int) MessageHandler{
	decorator := &BaseMessageDecorator{}
	switch level{
	case 1:
		return &EmojiMessageDecorator{decorator}
	
	case 2 :
		return &LinkMessageDecorator{decorator}

	case 3 :
		edecorator := &EmojiMessageDecorator{decorator}
		return &LinkMessageDecorator{edecorator}
	default:
		return decorator
	}
}

