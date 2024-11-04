package sender

import (
	"fmt"
	"misis/pkg/laba3/builder"
	Mlogger "misis/pkg/laba3/singleton"
	"misis/pkg/laba4/adapter"
	"misis/pkg/laba4/proxy"
)

type SenderFacade struct{
	auth *userproxy.UserProxy
 	message *builderMessage.MessageBuilder
}

func NewSenderFacade(logger *Mlogger.Mlogger , database database.Database) *SenderFacade{
	senderFacade:= &SenderFacade{
		auth : userproxy.NewUserProxy(logger , database),
		message: builderMessage.NewBuilder(),
	}

	return senderFacade
}

func (s *SenderFacade) Send(senderID int, destinationID int , message string , messageType string ){
	sender:=s.auth.GetData(senderID)
	des:= s.auth.GetData(destinationID)

	mes := s.message.SetSender(sender.Name).SetDestination(des.Name).SetMessage(messageType,message).Build()

	mes.Message.Create()
	fmt.Println(mes.Sender + " -> " + mes.Destination + " Message: "+ mes.Message.Get())
}
