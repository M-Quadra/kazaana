package kazaana

import (
	"fmt"
	"testing"
	"time"
)

func TestHelloWorld(t *testing.T) {
	fmt.Println("[kazaana] Error Printer Test......")

	kerr := errorFunc()
	kerr.HasError()
	kerr.HasError("[optional]:")
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
