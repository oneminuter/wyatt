package util

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

//日志
func Logger(str string, skips ...int) {
	skip := 1
	if len(skips) > 0 {
		skip = skips[0]
	}
	_, file, line, ok := runtime.Caller(skip)
	if ok {
		//runtime.FuncForPC(funcname).Name()
		now := time.Now().Format("2006-01-02 15:04:05")
		os.Stderr.Write([]byte(fmt.Sprintf("%s %s:%d %s \n", now, file, line, str)))
	}
}

//错误日志
func LoggerError(err error) {
	Logger(err.Error(), 2)
}

//调试日志
func LoggerDebug(err error) {
	os.Stderr.Write([]byte("\n" + time.Now().Format("2006-01-02 15:04:05") + " Error: " + err.Error() + "\n"))
	debug.PrintStack()
	os.Stderr.Write([]byte("\n"))
}
