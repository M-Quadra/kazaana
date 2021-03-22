package test

import (
	"fmt"
	"testing"

	"github.com/M-Quadra/kazaana"
)

func TestConfig(t *testing.T) {
	kazaana.Config.FirstCallers(3)
	kazaana.Config.Header("new header")
	fmt.Println(kazaana.Config.FirstCallers(), kazaana.Config.Header())
}
