package c_errors

type CError interface {
	Message() string
	Error() string
}
