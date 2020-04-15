package kazaana

import "time"

var (
	// FirstCallers first track num
	FirstCallers = 5

	// TimeLocation error information time
	//  default time.Local
	TimeLocation = time.Local

	// Header of error info
	//  default "error happen:"
	Header = "error happen:"
)

func init() {
	if TimeLocation != nil {
		return
	}

	TimeLocation = time.Local
}

func timeLocation() *time.Location {
	tl := TimeLocation
	if tl == nil {
		tl = time.Local
		TimeLocation = tl
	}
	return tl
}
