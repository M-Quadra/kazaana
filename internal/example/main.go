package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/M-Quadra/kazaana/v3"
)

func f0() error {
	return kazaana.New("err0", 0)
}

func f1() error {
	err1 := f0()
	if err1 != nil {
		return kazaana.Wrap(err1, 1)
	}

	// do something...
	return nil
}

func main() {
	if err := f1(); err != nil {
		PrintErr(err)
	}
}

// PrintErr print error info
func PrintErr(err error) {
	if err == nil {
		return
	}

	infos := [][]string{}
	lineCnt := 0

	for {
		if kerr, ok := err.(kazaana.Error); ok {
			arr := make([]string, 0, len(kerr.Callers)+2)
			arr = append(arr, "Time: "+kerr.Time.Format("2006-01-02 15:04:05.000"))
			if len(kerr.Args) > 0 {
				arr = append(arr, "Args: "+fmt.Sprint(kerr.Args...))
			}
			arr = append(arr, kerr.Callers...)

			infos = append(infos, arr)
			lineCnt += len(arr)
		} else {
			str := err.Error()
			if len(str) > 0 {
				infos = append(infos, []string{str})
				lineCnt++
			}
		}

		err = errors.Unwrap(err)
		if err == nil {
			break
		}
	}

	if lineCnt <= 0 {
		return
	}
	lines := make([]string, 0, lineCnt+1)
	lines = append(lines, "[Kazaana] error catch, "+time.Now().Format("2006-01-02 15:04:05"))

	prefix := ""
	for i := len(infos) - 1; i >= 0; i-- {
		for _, line := range infos[i] {
			if len(line) <= 0 {
				continue
			}
			lines = append(lines, prefix+line)
		}
		prefix += "    "
	}

	fmt.Println(strings.Join(lines, "\n"))
}
