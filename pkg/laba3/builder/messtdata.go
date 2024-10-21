package builderMessage

import (
	messt "misis/pkg/laba3/factoryMethod"
	."misis/pkg/laba3/singleton"
)

type MessageData struct{
	Sender string
	Destination string
	Message messt.Message
}



type MessageBuilder struct{
	message MessageData
}

func NewBuilder() *MessageBuilder{
	return &MessageBuilder{}
}

func (mb *MessageBuilder)SetSender(s string) *MessageBuilder{
	mb.message.Sender= s
	return mb
}

func (mb *MessageBuilder)SetDestination(d string)*MessageBuilder{
	mb.message.Destination = d
	return mb
}

func (mb *MessageBuilder)SetMessage(messageType string, message string)*MessageBuilder{
	logger := NewMlogger()
	logger.LevelDebug()
	mb.message.Message =  messt.MessageFactory(messageType,message,logger)
	return mb
}

func (mb *MessageBuilder)Build() *MessageData{
	return &mb.message
}