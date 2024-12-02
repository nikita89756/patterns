package messt

import (
	"fmt"
	photo "misis/pkg/laba3/abstractfactory"
	. "misis/pkg/laba3/singleton"
)
type Message interface {
	Create()
	Copy() Message
	Get()string
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

func (tm TextMessage) Create() {
 fmt.Println(tm.logger.LogError("Отправка текстового сообщения"))
}

func (tm TextMessage) Copy()Message{
	return TextMessage{tm.Message,tm.logger}

}

func (tm TextMessage) Get()string{
	return tm.Message

}

type VideoMessage struct {
	FilePath string
	logger *Mlogger
}

func (vm VideoMessage) Create() {
	fmt.Println(vm.logger.LogInfo("Отправка видео сообщения"))
}

func (vm VideoMessage) Copy()Message{
	return VideoMessage{vm.FilePath,vm.logger}

}

func (vm VideoMessage) Get()string{
	return vm.FilePath

}

type PhotoMessage struct {
	FilePath string
	logger *Mlogger
}

func (pm PhotoMessage) Create(){
	fmt.Println(pm.logger.LogInfo("Отправка фото сообщения"))
	photo.NewPngFactory(pm.logger).CreateImg(pm.FilePath)
}

func (pm PhotoMessage) Copy()Message{
	return PhotoMessage{pm.FilePath,pm.logger}

}

func (pm PhotoMessage) Get()string{
	return pm.FilePath

}

