package util

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Logger struct {
	buffer *bytes.Buffer
	log.Logger
}

func NewLogger(flag int) Logger {
	var buffer bytes.Buffer
	return Logger{&buffer, *log.New(&buffer, "", flag)}
}

func (l *Logger) Flush(image *ebiten.Image) {
	ebitenutil.DebugPrint(image, l.buffer.String())
	l.buffer.Reset()
}
