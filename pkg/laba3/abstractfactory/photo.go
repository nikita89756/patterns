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
	return "png -> gziped"
}

func (p *pngPhoto) GetDate() time.Time{
	return p.date
}

func (p *pngPhoto )GetSize() int{
	return p.size
}

type JpegFactory struct{
	logger *Mlogger
}

type jpegPhoto struct{
	date time.Time
	size int
}

func NewJpegFactory(log *Mlogger) PhotoFactory{
	return &JpegFactory{log}
}
func (png *JpegFactory) CreateImg(s string) Img{
	fmt.Println(png.logger.LogInfo("Фото: " +s))
	return &jpegPhoto{date:time.Now(),size: rand.Int()}
}

func (png *JpegFactory) CreateCompressed(s string)string{
	fmt.Println(png.logger.LogInfo("Фото: " +s+"сжато"))
	return "jpeg -> gziped"
}

func (p *jpegPhoto) GetDate() time.Time{
	return p.date
}

func (p *jpegPhoto )GetSize() int{
	return p.size
}


