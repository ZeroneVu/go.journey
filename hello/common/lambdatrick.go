package common

type Doer interface {
	Execute(...interface{})
}

type Executable func(...interface{})

// Satisfy Doer interface.
// So we can now pass an anonymous function using Executable,
// which implements Executable.
func (e Executable) Execute(args ...interface{}) {
	// delegate to the anonymous function
	e(args)
}
