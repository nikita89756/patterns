package photo

import (
	"fmt"
	"math/rand"
	. "misis/pkg/laba3/singleton"
	"time"
)


type PhotoFactory interface{
 CreateImg (string) Img
 CreateCompressed (string) string
}

type Img interface{
	GetDate() time.Time
	GetSize() int
}

type PngFactory struct{
	logger *Mlogger
}

type pngPhoto struct{
	date time.Time
	size int
}

func NewPngFactory(log *Mlogger) PhotoFactory{
	return &PngFactory{log}
}
func (png *PngFactory) CreateImg(s string) Img{
	fmt.Println(png.logger.LogInfo("Фото: " +s))
	return &pngPhoto{date:time.Now(),size: rand.Int()}
}

func (png *PngFactory) CreateCompressed(s string)string{
	fmt.Println(png.logger.LogInfo("Фото: " +s+"сжато"))
	return "gziped"
}

func (p *pngPhoto) GetDate() time.Time{
	return p.date
}

func (p *pngPhoto )GetSize() int{
	return p.size
}


