package kazaana

import (
	"fmt"
	"runtime"
	"time"
)

// FirstCallers first track num
var FirstCallers = 5

// Error error info
type Error struct {
	_beginTime time.Time
	_info      string
	_has       bool
	_callers   []string
}

// Empty check error and print
func (slf Error) Empty() bool {
	if !slf._has {
		return false
	}

	stTime := slf._beginTime

	fmt.Println("error happen:")
	fmt.Println("    ", stTime.Format("2006-01-02 15:04:05"), stTime.Unix(), stTime.UnixNano())
	fmt.Println("    ", slf._info)
	for _, v := range slf._callers {
		fmt.Println("     " + v)
	}
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
		_beginTime: time.Now(),
		_info:      err.Error(),
		_callers:   callerInfoAry,
		_has:       true,
	}
	return opt
}
