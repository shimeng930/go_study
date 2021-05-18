package log

var (
	_useIODiscard   = false
	_useJSONEncoder = false
)

func EnableJSONEncoder() (disable func()) {
	_useJSONEncoder = true
	return func() { _useJSONEncoder = false }
}

func EnableIODiscard() (disable func()) {
	_useIODiscard = true
	return func() { _useIODiscard = false }
}
