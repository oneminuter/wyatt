package util

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
)

/*
	日志:打印调用出错的位置
*/
func Logger(str string, skips ...int) {
	skip := 1
	if len(skips) > 0 {
		skip = skips[0]
	}
	_, file, line, ok := runtime.Caller(skip)
	if ok {
		//runtime.FuncForPC(funcname).Name()
		now := time.Now().Format("2006-01-02 15:04:05")
		os.Stderr.Write([]byte(fmt.Sprintf("%s %s:%d\t%s\n", now, file, line, str)))
	}
}

//打印多条信息
func Loggerln(str ...interface{}) {
	var outstr = make([]string, 0, len(str))
	for _, v := range str {
		switch v.(type) {
		case string:
			outstr = append(outstr, v.(string))
		default:
			bytes, err := jsoniter.Marshal(v)
			if err != nil {
				LoggerError(err)
			}
			outstr = append(outstr, string(bytes))
		}

	}
	Logger(strings.Join(outstr, " "), 2)
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
