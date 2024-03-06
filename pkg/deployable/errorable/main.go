package errorable

type ErrorCode int

const (
	ACTIONABLE_NOT_FOUND ErrorCode = iota
)

type Errorable interface {
	error
	Code() ErrorCode
}

type SErrorable struct {
	ErrorCode
	Message string
}

func NewErrorable(
	errorCode ErrorCode, message string,
) *SErrorable {
	return &SErrorable{ErrorCode: errorCode, Message: message}
}

func (e *SErrorable) Error() string {
	return e.Message
}

func (e *SErrorable) Code() ErrorCode {
	return e.ErrorCode
}
