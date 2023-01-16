package quickbooks

import (
	"fmt"
	"os"
	"strings"
)

var logger Loggative = &Log{level: 4}

type Log struct{ level int }
type Loggative interface {
	Err(string) // 4
	Errf(string, ...interface{})
	Warn(string) // 3
	Warnf(string, ...interface{})
	Info(string) // 2
	Infof(string, ...interface{})
	Debug(string) // 1
	Debugf(string, ...interface{})
}

func fmtLog(prefix, data string) {
	out := prefix
	lines := strings.Split(data, "\n")
	spaces := strings.Repeat(" ", len(prefix))
	for i, d := range lines {
		if i != 0 {
			out += "\n" + spaces
		}
		out += d
	}
	if len(out) > 0 && out[len(out)-1] != '\n' {
		out += "\n"
	}

	fmt.Fprintf(os.Stderr, out)
}

func (l *Log) Err(s string) {
	if l.level <= 4 {
		fmtLog("[ERR ]: ", s)
	}
}
func (l *Log) Errf(f string, is ...interface{}) {
	if l.level <= 4 {
		fmtLog("[ERR ]: ", fmt.Sprintf(f, is...))
	}
}
func (l *Log) Warn(s string) {
	if l.level <= 3 {
		fmtLog("[WRN]: ", s)
	}
}
func (l *Log) Warnf(f string, is ...interface{}) {
	if l.level <= 3 {
		fmtLog("[WRN]: ", fmt.Sprintf(f, is...))
	}
}
func (l *Log) Info(s string) {
	if l.level <= 2 {
		fmtLog("[INF]: ", s)
	}
}
func (l *Log) Infof(f string, is ...interface{}) {
	if l.level <= 2 {
		fmtLog("[INF]: ", fmt.Sprintf(f, is...))
	}
}
func (l *Log) Debug(s string) {
	if l.level <= 1 {
		fmtLog("[DBG]: ", s)
	}
}
func (l *Log) Debugf(f string, is ...interface{}) {
	if l.level <= 1 {
		fmtLog("[DBG]: ", fmt.Sprintf(f, is...))
	}
}

func (l *Log) LogLevel(i int) {
	l.level = i
}

func SetLogger(l *Loggative) {
	logger = l
}
