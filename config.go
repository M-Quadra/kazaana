package kazaana

// Config global config
var Config = struct {
	Creator Creatable
	Wrapper Wrapable
}{
	Creator: defaultCreator{},
	Wrapper: defaultWrapper{},
}

// Unwrapable error+Unwrap
type Unwrapable interface {
	Unwrap() error
}
