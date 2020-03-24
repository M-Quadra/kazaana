package kazaana

import "time"

// FirstCallers first track num
var FirstCallers = 5

// TimeLocation error information time
// default time.Local
var TimeLocation *time.Location

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
