package kazaana

import (
	"fmt"
	"runtime"
	"time"
)

// Error error info
type Error struct {
	_beginTime time.Time
	_callers   []string
	_src       error
}

// HasError error check
func HasError(err error) bool {
	if err == nil {
		return false
	}

	nowTime := time.Now().In(timeLocation())

	info := fmt.Sprintln(Header)
	info += fmt.Sprintln("    ", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix(), nowTime.UnixNano())
	info += fmt.Sprintln("    ", err.Error())
	for i := 1; i < FirstCallers+1; i++ {
		ptr, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}

		info += fmt.Sprintf("     %s:%d +%#x", file, line, ptr) + "\n"
	}
	fmt.Println(info)
	return true
}

// CheckError check error without print
func (slf Error) CheckError() bool {
	return slf._src != nil
}

// HasError check error and print
func (slf Error) HasError() bool {
	if !slf.CheckError() {
		return false
	}

	stTime := slf._beginTime

	info := fmt.Sprintln(Header)
	info += fmt.Sprintln("    ", stTime.Format("2006-01-02 15:04:05"), stTime.Unix(), stTime.UnixNano())
	info += fmt.Sprintln("    ", slf._src.Error())
	for _, v := range slf._callers {
		info += fmt.Sprintln("    ", v)
	}
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
		_beginTime: time.Now().In(timeLocation()),
		_callers:   callerInfoAry,
		_src:       err,
	}
	return opt
}

// RawError get raw error
func (slf Error) RawError() error {
	return slf._src
}
