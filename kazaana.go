package kazaana

import "errors"

// New create a new error with string
func New(str string, args ...any) error {
	creator := Config.Creator
	if len(str) <= 0 || creator == nil {
		return nil
	}

	return creator.Create(errors.New(str), args...)
}

// Wrap wrap a error
func Wrap(err error, args ...any) error {
	wrapper := Config.Wrapper
	if err == nil || wrapper == nil {
		return err
	}

	return wrapper.Wrap(err, args...)
}
