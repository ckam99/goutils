package log

import (
	"encoding/json"
	"fmt"
	baseLog "log"
	"os"
	"sync"
	"time"
)

type Level string

const (
	ERROR Level = "error"
	INFO  Level = "info"
	WARN  Level = "warning"
	DEBUG Level = "debug"
)

type Format string

const (
	JSON    Format = "json"
	DEFAULT Format = ""
)

var LogFormat = DEFAULT

type Log struct {
	mutex       sync.RWMutex
	LogLevel    Level         `json:"level"`
	LogProtocol string        `json:"protocol"`
	LogMethod   string        `json:"method"`
	StatusCode  int           `json:"code"`
	Dur         time.Duration `json:"duration"`
	Message     any           `json:"message"`
	Time        string        `json:"time"`
	//format      Format
	outputPath string
}

var instance *Log

func setInstanceLevel(level Level) *Log {
	if instance == nil {
		instance = &Log{}
	}
	instance.LogLevel = level
	return instance
}

func GetFile() (*os.File, error) {
	if instance != nil {
		return instance.GetFile(), nil
	}
	return nil, fmt.Errorf("logs instance is not initialized")
}

func Info() *Log {
	return setInstanceLevel(INFO)
}

func Debug() *Log {
	return setInstanceLevel(DEBUG)
}

func Error() *Log {
	return setInstanceLevel(ERROR)
}

func Warn() *Log {
	return setInstanceLevel(WARN)
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

//
//func (l *Log) Format(format Format) *Log {
//	//l.format = format
//	return l
//}

func (l *Log) Output(filePath string) *Log {
	l.outputPath = filePath
	return l
}

func (l *Log) Fatal(message any) {
	l.Msg("%s", message)
	baseLog.Fatal(message)
}

func (l *Log) GetFile() *os.File {
	// l.format = Format(os.Getenv("LOG_FORMAT"))

	if l.outputPath == "" {
		l.outputPath = os.Getenv("LOG_FILE")
		if l.outputPath == "" {
			l.outputPath = "logs.log"
		}
	}
	f, err := os.OpenFile(l.outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		baseLog.Printf("logging: %s", err.Error())
		return nil
	}
	return f
}

func (l *Log) Str(format string, message ...any) {
	l.mutex.RLock()
	l.Message = fmt.Sprintf(format, message...)
	f := l.GetFile()
	defer f.Close()

	lg := baseLog.New(f, "", baseLog.LstdFlags)
	output := fmt.Sprintf("[%v]\t%s\t%s\t%d\t%s\t%v", l.LogLevel, l.LogProtocol, l.LogMethod, l.StatusCode, l.Dur, l.Message)
	baseLog.Println(output)
	lg.Println(output)
	l.mutex.RUnlock()
}

func (l *Log) Json(format string, message ...any) {
	l.mutex.RLock()
	l.Message = fmt.Sprintf(format, message...)
	f := l.GetFile()
	defer f.Close()
	l.Time = time.Now().String()
	b, _ := json.Marshal(l)
	fmt.Println(string(b))
	f.WriteString(string(b) + "\n")
	l.mutex.RUnlock()
}

func (l *Log) Msg(format string, message ...any) {

	//if l.format == JSON {
	if LogFormat == JSON {
		l.Json(format, message...)
	} else {
		l.Str(format, message...)
	}
}
