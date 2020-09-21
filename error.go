package kazaana

import (
	"fmt"
	"runtime"
	"time"
)

// Error error info
type Error struct {
	beginTime time.Time
	callers   []string
	src       error
}

// HasError error check
func HasError(err error, header ...string) bool {
	if err == nil {
		return false
	}

	callerInfoAry := []string{}
	for i := 1; i < FirstCallers+1; i++ {
		ptr, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		info := fmt.Sprintf("%s:%d +%#x", file, line, ptr)
		callerInfoAry = append(callerInfoAry, info)
	}

	nowTime := time.Now()
	nowHeader := Header
	if len(header) > 0 {
		nowHeader = header[0]
	}

	info := errorInfo(nowHeader, nowTime, err.Error(), callerInfoAry)
	fmt.Println(info)
	return true
}

// CheckError check error without print
func (slf Error) CheckError() bool {
	return slf.src != nil
}

// HasError check error and print
func (slf Error) HasError(header ...string) bool {
	if !slf.CheckError() {
		return false
	}

	stTime := slf.beginTime
	nowHeader := Header
	if len(header) > 0 {
		nowHeader = header[0]
	}

	info := errorInfo(nowHeader, stTime, slf.src.Error(), slf.callers)
	fmt.Println(info)
	return true
}

// New error to kazaana.Error
func New(err error) Error {
	if err == nil {
		return Error{}
	}

	callerInfoAry := []string{}
	for i := 1; i < FirstCallers+1; i++ {
		ptr, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		info := fmt.Sprintf("%s:%d +%#x", file, line, ptr)
		callerInfoAry = append(callerInfoAry, info)
	}

	opt := Error{
		beginTime: time.Now(),
		callers:   callerInfoAry,
		src:       err,
	}
	return opt
}

// RawError get raw error
func (slf Error) RawError() error {
	return slf.src
}

func errorInfo(header string, stTime time.Time, errStr string, callerAry []string) string {
	info := fmt.Sprintln(header)
	info += fmt.Sprintln("    ", stTime.Format("2006-01-02 15:04:05"), stTime.UnixNano())
	info += fmt.Sprintln("    ", errStr)
	for _, v := range callerAry {
		info += fmt.Sprintln("    ", v)
	}

	return info
}
