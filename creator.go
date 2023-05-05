package kazaana

import (
	"fmt"
	"runtime"
	"time"
)

// Creatable custom error creator
type Creatable interface {
	Create(err error, args ...any) Error
}

type defaultCreator struct{}

var ( // check interface
	_ Creatable = defaultCreator{}
)

func (c defaultCreator) Create(err error, args ...any) Error {
	callers := make([]string, 0, 5)
	for i := 0; i < 5; i++ {
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
