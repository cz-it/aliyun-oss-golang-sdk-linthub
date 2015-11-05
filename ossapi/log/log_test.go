/**
* Author: CZ cz.theng@gmail.com
 */

package log

import (
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	l, err := NewFileLogger(".ossapilog", "ossapi")
	if err != nil {
		fmt.Errorf("create Logger Error\n")
		return
	}
	l.SetMaxFileSize(10000) //100MB
	l.SetLevel(LDEBUG)
	l.Error("errorErrorError")
	l.Debug("Debug:int a is %d", 10)
	l.Fatal("fatal")
	l.Info("Info")
	l.SetCallDepth(2)
	l.Warning("Warning")
	l.Info("", err)
	l.SetLevel(LFATAL)
	l.Error("error")
	l.Debug("Debug:int a is %d", 10)
	l.Fatal("fatal")
	l.Info("Info")
	l.SetCallDepth(2)
	l.Warning("Warning")
	l.Info("", err)

	DEBUG("debug")
	INFO("info")
	WARNING("warning")
	ERROR("error")
	FATAL("fatal")
}
func TestConsoleLog(t *testing.T) {
	l, err := NewConsoleLogger()
	if err != nil {
		fmt.Errorf("create Logger Error\n")
		return
	}
	l.Info("aa")
	l.Error("bb")
}

func TestLogFileLog_2(t *testing.T) {
	l, err := NewFileLogger("/dev/a", "txt")
	if err != nil {
		fmt.Errorf("create Logger Error\n")
		return
	}
	l.SetCallDepth(1024)
	l.Debug("Debug")
	l.Fatal("fatal")
}

func TestLogSizeFile(t *testing.T) {
	l, err := NewFileLogger(".log", "ossapi")
	if err != nil {
		fmt.Errorf("create Logger Error\n")
		return
	}
	l.SetMaxFileSize(200) //100MB
	l.SetLevel(LDEBUG)
	l.Info("1234567890")
	l.Info("1234567890")
	l.Info("1234567890")
	l.Info("1234567890")
	l.Info("1234567890")
}
