package kazaana

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

// Wrapable custom error wrapper
type Wrapable interface {
	Wrap(err error, args ...any) Error
}

type defaultWrapper struct{}

var ( // check interface
	_ Wrapable = defaultWrapper{}
)

func (w defaultWrapper) Wrap(err error, args ...any) Error {
	maxCnt := 5
	var kerr *Error
	if errors.As(err, &kerr) {
		maxCnt = 1
	}

	callers := make([]string, 0, maxCnt)
	for i := 0; i < maxCnt; i++ {
		pc, file, line, ok := runtime.Caller(2 + i)
		if !ok {
			break
		} else if len(file) <= 0 {
			continue
		}

		f := runtime.FuncForPC(pc)
		info := fmt.Sprintf("%s:%d +%s", file, line, f.Name())
		callers = append(callers, info)
	}

	return Error{
		Raw:     err,
		Time:    time.Now(),
		Args:    args,
		Callers: callers,
	}
}
