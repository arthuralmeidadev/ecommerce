package api

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type Logger struct {
	Context string
}

var onceLogger sync.Once
var singleLogger *Logger

func GetLogger() *Logger {
	onceLogger.Do(func() {
		singleLogger = &Logger{
			Context: "Application",
		}
	})

	return singleLogger
}

func (l *Logger) Info(msg string) {
	l.print("\x1b[90m", msg, "INFO")
}

func (l *Logger) Debug(msg string) {
	l.print("\x1b[95m", msg, "DEBUG")
}

func (l *Logger) Warn(msg string) {
	l.print("\x1b[93m", msg, "WARNING")
}

func (l *Logger) Success(msg string) {
	l.print("\x1b[92m", msg, "SUCCESS")
}

func (l *Logger) Error(msg string) {
	l.print("\x1b[91m", msg, "ERROR")
}

func (l *Logger) Fatal(msg string) {
	l.print("\x1b[91m", msg, "FATAL")
	os.Exit(1)
}

func (l *Logger) print(color, msg, msgType string) {
	t := time.Now()
	b := &strings.Builder{}
	b.WriteString("\x1b[97m[\x1b[93m")
	b.WriteString(t.Format(time.DateOnly))

	var h, period string
	m := fmt.Sprintf("%d", t.Minute())
	s := fmt.Sprintf("%d", t.Second())

	period = "AM"
	if t.Hour() == 0 {
		h = "12"
	} else if t.Hour() > 12 {
		h = fmt.Sprintf("%d", t.Hour()-12)
		period = "PM"
	}

	if len(h) < 2 {
		h = fmt.Sprintf("0%s", h)
	}

	if len(m) < 2 {
		m = fmt.Sprintf("0%s", m)
	}

	if len(s) < 2 {
		s = fmt.Sprintf("0%s", s)
	}

	b.WriteString(fmt.Sprintf(" %s:%s:%s %s", h, m, s, period))
	b.WriteString("\x1b[97m] [\x1b[93m")
	b.WriteString(l.Context)
	b.WriteString("\x1b[97m]\x1b[0m ")
	prefix := b.String()
	if len(prefix) < 74 {
		b.WriteString(strings.Repeat(" ", 74-len(prefix)))
	}

	b.WriteString(strings.Replace(color, "9", "4", 1))
	if msgType == "INFO" {
		b.WriteString("\x1b[37m\x1b[1m ")
	} else {
		b.WriteString("\x1b[38;2;0;0;0m\x1b[1m ")
	}
	b.WriteString(msgType)
	b.WriteString(":")
	if len(msgType) < 9 {
		b.WriteString(strings.Repeat(" ", 9-len(msgType)))
	}
	b.WriteString("\x1b[0m")
	b.WriteString(" ")
	b.WriteString(color)
	b.WriteString(msg)
	b.WriteString("\x1b[0m\n")
	fmt.Print(b.String())
}

func launch(msg string) {
	t := time.Now()
	b := &strings.Builder{}
	b.WriteString("\x1b[97m[\x1b[93m")
	b.WriteString(t.Format(time.DateOnly))

	var h, period string
	m := fmt.Sprintf("%d", t.Minute())
	s := fmt.Sprintf("%d", t.Second())

	period = "AM"
	if t.Hour() == 0 {
		h = "12"
	} else if t.Hour() > 12 {
		h = fmt.Sprintf("%d", t.Hour()-12)
		period = "PM"
	}

	if len(h) < 2 {
		h = fmt.Sprintf("0%s", h)
	}

	if len(m) < 2 {
		m = fmt.Sprintf("0%s", m)
	}

	if len(s) < 2 {
		s = fmt.Sprintf("0%s", s)
	}

	b.WriteString(fmt.Sprintf(" %s:%s:%s %s", h, m, s, period))
	b.WriteString("\x1b[97m] [\x1b[93m")
	b.WriteString("Application")
	b.WriteString("\x1b[97m]\x1b[0m ")
	prefix := b.String()
	if len(prefix) < 74 {
		b.WriteString(strings.Repeat(" ", 74-len(prefix)))
	}
	b.WriteString("\x1b[46m\x1b[38;2;0;0;0m\x1b[1m LAUNCHED: ")
	b.WriteString("\x1b[0m")
	b.WriteString(" ")
	b.WriteString("\x1b[96m")
	b.WriteString(msg)
	b.WriteString("\x1b[0m\n")
	fmt.Print(b.String())
}
