package kazaana

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestHelloWorld(t *testing.T) {
	fmt.Println("[kazaana] Error Printer Test......")

	kerr := errorFunc()
	kerr.HasError()
	kerr.HasError("[optional]:")

	err := errors.New("WTF")
	fmt.Println(err)
	err = &kerr
	fmt.Println(err)
	fmt.Println(kerr)

	kerr = New(nil)
	fmt.Println(kerr.Error())
}

func errorFunc() Error {
	kerr := errorHappen()
	return kerr
}

func errorHappen() Error {
	timeStr := "1970-01-01 08:00:00"
	_, err := time.Parse("2006-01-02 15:04:051", timeStr)
	HasError(err)
	HasError(err, "[optional2]:")
	return New(err)
}
