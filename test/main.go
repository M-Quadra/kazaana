package main

import (
	"fmt"
	"time"

	"github.com/m_quadra/kazaana"
)

func main() {
	kerr := errorTest2()

	if kerr.HasError() {
		fmt.Println("err")
	} else {
		fmt.Println("ok")
	}
}

func errorTest2() kazaana.Error {
	kerr := errorTest()
	return kerr
}

func errorTest() kazaana.Error {
	fmt.Println("\nError Printer Test......")

	timeStr := "1970-01-01 08:00:00"
	_, err := time.Parse("2006-01-02 15:04:051", timeStr)
	kazaana.HasError(err)
	kerr := kazaana.New(err)
	return kerr
}
