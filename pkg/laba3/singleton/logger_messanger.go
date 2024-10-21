package Mlogger

import (
	"fmt"
	"sync"
	"time"
)

type Mlogger struct{
	level int
	substring string
}

	var (
		once sync.Once
		instance *Mlogger
	)
func NewMlogger() *Mlogger{
	
	once.Do(func(){instance = &Mlogger{0,""}})
	return instance
}

func (m *Mlogger) LevelDebug(){
	m.level = 0
}

func (m *Mlogger) LevelInfo(){
	m.level = 1
}

func (m *Mlogger) LogInfo(str string)string{
	if m.level ==0{
		return fmt.Sprintf("%s :: %s",time.Now(),str)
	}
	return ""
}

func (m *Mlogger) LogError(str string)string{
	if m.level >=0{
		return fmt.Sprintf("%s :: %s",time.Now(),str)
	}
	return ""
}