package sender

import (
	"fmt"
	"misis/pkg/laba3/builder"
	Mlogger "misis/pkg/laba3/singleton"
	"misis/pkg/laba4/adapter"
	messagehandler "misis/pkg/laba4/decorator"
	"misis/pkg/laba4/proxy"
)

type SenderFacade struct{
	auth *userproxy.UserProxy
	messageDecorator messagehandler.MessageHandler
 	message *builderMessage.MessageBuilder
}

func NewSenderFacade(logger *Mlogger.Mlogger , database database.Database , level int) *SenderFacade{
	senderFacade:= &SenderFacade{
		auth : userproxy.NewUserProxy(logger , database),
		messageDecorator: messagehandler.CreateMessageHandler(level),
		message: builderMessage.NewBuilder(),
	}

	return senderFacade
}

func (s *SenderFacade) Send(senderID int, destinationID int , message string , messageType string ){
	sender:=s.auth.GetData(senderID)
	des:= s.auth.GetData(destinationID)

	s.messageDecorator.ProcessMessage(message)
	message = s.messageDecorator.GetString()

	mes := s.message.SetSender(sender.Name).SetDestination(des.Name).SetMessage(messageType,message).Build()

	mes.Message.Create()
	fmt.Println(mes.Sender + " -> " + mes.Destination + " Message: "+ mes.Message.Get())
}
