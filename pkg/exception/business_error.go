package exception

type BusinessError struct {
	message string
}

func (e BusinessError) Error() string {
	return e.message
}

func NewBusinessError(message string) error {
	return &BusinessError{message: message}
}
