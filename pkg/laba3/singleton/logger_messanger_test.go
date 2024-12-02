package Mlogger

import "testing"

func TestNewMlogger(t *testing.T) {
	ml:=NewMlogger()
	l:=NewMlogger()
	if ml != l{
		t.Error("Ошибка дублирования логгеров")
	}
}

func TestLevelInfo(t *testing.T){
	l:=NewMlogger()
	l.LevelInfo()
	if l.LogInfo("logger")!=""{
		t.Error("Ошибка логгер вернул текс")
	}
	if l.LogError("logger")==""{
		t.Error("Ошибка логгер должен вернуть текст")
	}
}
func TestLevelDebug(t *testing.T){
	l:=NewMlogger()
	l.LevelDebug()
	if l.LogInfo("logger")==""{
		t.Error("Ошибка логгер должен вернуть текст")
	}
	if l.LogError("logger")==""{
		t.Error("Ошибка логгер должен вернуть текст")
	}
}