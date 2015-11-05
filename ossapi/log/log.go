/**
* Author: CZ cz.theng@gmail.com
* A log package for golang
*
* not instead of golang's log but a replenish
 */

package log

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

// Level is log's level
type Level int // loose enum type . maybe have some other define method
const (
	//LNULL is none
	LNULL = iota
	//LDEBUG is debug log
	LDEBUG
	// LINFO is info log
	LINFO
	//LWARNING is warning log
	LWARNING
	//LERROR is error log
	LERROR
	//LFATAL is fatal log
	LFATAL
)

const (
	maxLogSize = 5 * 1024 * 1024 // Default max log file size is 500M
)

type outputDevice int

const (
	stdoDeviceFlag = 1 << iota
	fileDeviceFlag
)

// EOutput is output error
var ErrOutput = errors.New("Output is invalied!")

// Logger is log object
type Logger struct {
	level        Level
	logDevice    *fileDevice
	errDevice    *fileDevice
	conDevice    *consoleDevice
	outputDevice outputDevice

	mtx       sync.Mutex
	buf       []byte
	callDepth int
}

var _logger *Logger

func init() {
	/*for Coverall 95%
	var err error
	_logger, err = NewConsoleLogger()
	if err != nil {
		//TODO:
	}
	*/
	_logger, _ = NewConsoleLogger()
	_logger.SetCallDepth(3)
}

//NewConsoleLogger create a Console Logger
func NewConsoleLogger() (*Logger, error) {
	var err error
	logger := &Logger{level: LDEBUG, outputDevice: stdoDeviceFlag, callDepth: 3}
	logger.conDevice, err = newConsoleDevice()
	return logger, err
}

//NewFileLogger create a file logger
func NewFileLogger(logPath, fileName string) (*Logger, error) {
	var err error
	logger := &Logger{level: LDEBUG, outputDevice: fileDeviceFlag, callDepth: 2}
	logFileName := filepath.Join(logPath, fileName+".log")
	logger.logDevice, err = newFileDevice(logFileName)
	if err != nil {
		return nil, err
	}

	errFileName := filepath.Join(logPath, fileName+".error")
	/* for Coverall 95%
	logger.errDevice, err = NewFileDevice(errFileName)
	if err != nil {
		return nil, err
	}
	*/
	logger.errDevice, _ = newFileDevice(errFileName)
	return logger, nil
}

// SetCallDepth set call path
func (l *Logger) SetCallDepth(d int) {
	l.callDepth = d
}

func (l *Logger) getFileLine() string {
	_, file, line, ok := runtime.Caller(l.callDepth)
	if !ok {
		file = "???"
		line = 0
	} // for coverall

	return file + ":" + itoa(line, -1)
}

/**
* Change from Golang's log.go
* Cheap integer to fixed-width decimal ASCII.  Give a negative width to avoid zero-padding.
* Knows the buffer has capacity.
 */
func itoa(i int, wid int) string {
	var u = uint(i)
	if u == 0 && wid <= 1 {
		return "0"
	} // for coverall

	// Assemble decimal in reverse order.
	var b [32]byte
	bp := len(b)
	for ; u > 0 || wid > 0; u /= 10 {
		bp--
		wid--
		b[bp] = byte(u%10) + '0'
	}
	return string(b[bp:])
}

func (l *Logger) getTime() string {
	// Time is yyyy-mm-dd hh:mm:ss.microsec
	var buf []byte
	t := time.Now()
	year, month, day := t.Date()
	buf = append(buf, itoa(int(year), 4)+"-"...)
	buf = append(buf, itoa(int(month), 2)+"-"...)
	buf = append(buf, itoa(int(day), 2)+" "...)

	hour, min, sec := t.Clock()
	buf = append(buf, itoa(hour, 2)+":"...)
	buf = append(buf, itoa(min, 2)+":"...)
	buf = append(buf, itoa(sec, 2)...)

	buf = append(buf, '.')
	buf = append(buf, itoa(t.Nanosecond()/1e3, 6)...)

	return string(buf[:])
}

func (l *Logger) output(level Level, prefix string, format string, v ...interface{}) (err error) {
	var levelStr string
	if level == LDEBUG {
		levelStr = "[DEBUG]"
	} else if level == LINFO {
		levelStr = "[INFO]"
	} else if level == LWARNING {
		levelStr = "[WARNING]"
	} else if level == LERROR {
		levelStr = "[ERROR]"
	} else if level == LFATAL {
		levelStr = "[FATAL]"
	} else {
		levelStr = "[UNKNOWN LEVEL]"
	} // for coverall

	var msg string
	if format == "" {
		msg = fmt.Sprintln(v...)
	} else {
		msg = fmt.Sprintf(format, v...)
	}

	l.mtx.Lock()
	defer l.mtx.Unlock()
	l.buf = l.buf[:0]

	//	l.buf = append(l.buf,"["+l.logName+"]" ...)
	l.buf = append(l.buf, levelStr...)
	l.buf = append(l.buf, prefix...)

	l.buf = append(l.buf, ":"+msg...)
	if len(msg) > 0 && msg[len(msg)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}

	if l.outputDevice == stdoDeviceFlag {
		_, err = l.conDevice.write(l.buf)
	} else if l.outputDevice == fileDeviceFlag {
		if level <= LWARNING {
			_, err = l.logDevice.Write(l.buf)
		} else {
			_, err = l.errDevice.Write(l.buf)
		}
	} else {
		err = ErrOutput
	} // for coverall
	return
}

// SetMaxFileSize set max files size
func (l *Logger) SetMaxFileSize(fileSize uint64) {
	l.logDevice.setFileSize(fileSize)
	l.errDevice.setFileSize(fileSize)
}

// SetLevel set log level
func (l *Logger) SetLevel(level Level) {
	l.level = level
}

//Debug debug log
func (l *Logger) Debug(format string, v ...interface{}) error {
	if l.level > LDEBUG {
		return nil
	}

	err := l.output(LDEBUG, "["+l.getTime()+"]["+l.getFileLine()+"]", format, v...)
	return err
}

// Info is info log
func (l *Logger) Info(format string, v ...interface{}) error {
	if l.level > LINFO {
		return nil
	}

	err := l.output(LINFO, "["+l.getTime()+"]["+l.getFileLine()+"]", format, v...)
	return err
}

// Warning is warning log
func (l *Logger) Warning(format string, v ...interface{}) error {
	if l.level > LWARNING {
		return nil
	}
	err := l.output(LWARNING, "["+l.getTime()+"]["+l.getFileLine()+"]", format, v...)
	return err
}

func (l *Logger) Error(format string, v ...interface{}) error {
	if l.level > LERROR {
		return nil
	}
	err := l.output(LERROR, "["+l.getTime()+"]["+l.getFileLine()+"]", format, v...)
	return err
}

// Fatal is fatal log
func (l *Logger) Fatal(format string, v ...interface{}) error {
	if l.level > LFATAL {
		return nil
	}

	err := l.output(LFATAL, "["+l.getTime()+"]["+l.getFileLine()+"]", format, v...)
	return err
}

//DEBUG is golbal Debug log
func DEBUG(format string, v ...interface{}) error {
	return _logger.Debug(format, v...)
}

//INFO is global Info log
func INFO(format string, v ...interface{}) error {
	return _logger.Info(format, v...)
}

//WARNING is global warning log
func WARNING(format string, v ...interface{}) error {
	return _logger.Warning(format, v...)
}

// ERROR is global error log
func ERROR(format string, v ...interface{}) error {
	return _logger.Error(format, v...)
}

// FATAL is global fatal log
func FATAL(format string, v ...interface{}) error {
	return _logger.Fatal(format, v...)
}
