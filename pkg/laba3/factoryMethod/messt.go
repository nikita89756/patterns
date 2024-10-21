package messt

import (
	"fmt"
	photo "misis/pkg/laba3/abstractfactory"
	. "misis/pkg/laba3/singleton"
)
type Message interface {
	Send()
	Copy() Message
}

func MessageFactory(messageType string, message string , logger *Mlogger) Message {
	switch messageType {
	case "text":
		return TextMessage{Message: message,logger:logger}

	case "video":
		return VideoMessage{FilePath: message,logger:logger}

	case "photo":
		return PhotoMessage{FilePath: message,logger:logger}

	default:
		return nil
	}
}

type TextMessage struct {
	Message string
	logger *Mlogger
}

func (tm TextMessage) Send() {
 fmt.Println(tm.logger.LogError("Отправка текстового сообщения"))
}

func (tm TextMessage) Copy()Message{
	return TextMessage{tm.Message,tm.logger}

}

type VideoMessage struct {
	FilePath string
	logger *Mlogger
}

func (vm VideoMessage) Send() {
	fmt.Println(vm.logger.LogInfo("Отправка видео сообщения"))
}

func (vm VideoMessage) Copy()Message{
	return TextMessage{vm.FilePath,vm.logger}

}

type PhotoMessage struct {
	FilePath string
	logger *Mlogger
}

func (pm PhotoMessage) Send(){
	fmt.Println(pm.logger.LogInfo("Отправка фото сообщения"))
	photo.NewPngFactory(pm.logger).CreateImg(pm.FilePath)
}

func (pm PhotoMessage) Copy()Message{
	return TextMessage{pm.FilePath,pm.logger}

}


