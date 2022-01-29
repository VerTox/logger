package logger

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Logger struct {
	TraceId string
}

func New(traceId string) *Logger {
	if traceId == "" {
		traceId = uuid.New().String()
	}

	return &Logger{
		TraceId: traceId,
	}
}

type Log struct {
	TraceId string      `json:"trace_id"`
	Level   string      `json:"level"`
	Message string      `json:"message"`
	Date    time.Time   `json:"date"`
	Payload interface{} `json:"payload,omitempty"`
}

func send(traceId, message, level string, payload interface{}) {
	log, _ := json.Marshal(&Log{
		TraceId: traceId,
		Message: message,
		Level:   level,
		Date:    time.Now(),
		Payload: payload,
	})

	fmt.Printf("%s\n", string(log))
}

func (l *Logger) Debug(message string, payload interface{}) {
	send(l.TraceId, message, "Debug", payload)
}

func (l *Logger) Info(message string, payload interface{}) {
	send(l.TraceId, message, "Info", payload)
}

func (l *Logger) Notice(message string, payload interface{}) {
	send(l.TraceId, message, "Notice", payload)
}

func (l *Logger) Warning(message string, payload interface{}) {
	send(l.TraceId, message, "Warning", payload)
}

func (l *Logger) Error(message string, payload interface{}) {
	send(l.TraceId, message, "Error", payload)
}

func (l *Logger) Critical(message string, payload interface{}) {
	send(l.TraceId, message, "Critical", payload)
}

func (l *Logger) Alert(message string, payload interface{}) {
	send(l.TraceId, message, "Alert", payload)
}
func (l *Logger) Emergency(message string, payload interface{}) {
	send(l.TraceId, message, "Emergency", payload)
}

func (l *Logger) Trace() string {
	return l.TraceId
}
