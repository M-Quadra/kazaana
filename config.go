package kazaana

import "sync"

var (
	rw = sync.RWMutex{}

	firstCallers = 5
	header       = "error happen:"
)

type config int

const (
	// Config config about kazaana
	Config config = 0
)

// FirstCallers first track n
//  get: len(n) <= 0
//  set: n[0]
//  default: 5
func (slf config) FirstCallers(n ...int) int {
	if len(n) <= 0 {
		rw.RLock()
		defer rw.RUnlock()

		return firstCallers
	}

	rw.Lock()
	defer rw.Unlock()

	firstCallers = n[0]
	return n[0]
}

// Header of error info
//  get: len(h) <= 0
//  set: h[0]
//  default: "error happen:"
func (slf config) Header(h ...string) string {
	if len(h) <= 0 {
		rw.RLock()
		defer rw.RUnlock()

		return header
	}

	rw.Lock()
	defer rw.Unlock()

	header = h[0]
	return h[0]
}
