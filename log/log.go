package log

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/ckam225/goutils/env"
)

type Level string

const (
	ERROR Level = "error"
	INFO  Level = "info"
	WARN  Level = "warning"
)

type Format string

const (
	JSON    Format = "json"
	DEFAULT Format = "default"
)

type Log struct {
	LogLevel    Level         `json:"level"`
	LogProtocol string        `json:"protocol"`
	LogMethod   string        `json:"method"`
	StatusCode  int           `json:"code"`
	Dur         time.Duration `json:"duration"`
	Time        string        `json:"time"`
	format      Format
	outputPath  string
}

func Info() *Log {
	return &Log{
		LogLevel: INFO,
		format:   DEFAULT,
	}
}

func Error() *Log {
	return &Log{
		LogLevel: ERROR,
		format:   DEFAULT,
	}
}

func Warn() *Log {
	return &Log{
		LogLevel: WARN,
		format:   DEFAULT,
	}
}

func (l *Log) Protocol(protocol string) *Log {
	l.LogProtocol = protocol
	return l
}

func (l *Log) Method(method string) *Log {
	l.LogMethod = method
	return l
}

func (l *Log) Code(code int) *Log {
	l.StatusCode = code
	return l
}

func (l *Log) Duration(duration time.Duration) *Log {
	l.Dur = duration
	return l
}

func (l *Log) Format(format Format) *Log {
	l.format = format
	return l
}

func (l *Log) Output(filePath string) *Log {
	l.outputPath = filePath
	return l
}

func (l *Log) Msg(message string) {
	if l.outputPath == "" {
		l.outputPath = filepath.Join(env.GetString("LOG_DIR", "./"), "logs.log")
	}
	f, err := os.OpenFile(l.outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("logging: %s", err.Error())
		return
	}
	defer f.Close()

	if l.format == JSON {
		l.Time = time.Now().String()
		b, _ := json.Marshal(l)
		fmt.Println(string(b))
		f.WriteString(string(b) + "\n")
	} else {
		lg := log.New(f, "", log.LstdFlags)
		output := fmt.Sprintf("[%v]\t%s\t%s\t%d\t%s\t%s\n", l.LogLevel, l.LogProtocol, l.LogMethod, l.StatusCode, l.Dur, message)
		log.Println(output)
		lg.Println(output)
	}
}
