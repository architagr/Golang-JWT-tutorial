package models

const (
	ErrorTypeFatal        = "Fatal"
	ErrorTypeError        = "Error"
	ErrorTypeValidation   = "Validation Error"
	ErrorTypeInfo         = "Info"
	ErrorTypeDebug        = "Debug"
	ErrorTypeUnauthorized = "Unauthorized"
)

type ErrorDetail struct {
	ErrorType    string
	ErrorMessage string
}
type Response struct {
	Data    interface{}
	Status  int
	Error   []ErrorDetail
	Message string
}