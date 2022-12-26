package exception

type InfrastructureError struct {
	message string
}

func (e InfrastructureError) Error() string {
	return e.message
}

func NewInfrastructureError(message string) error {
	return &InfrastructureError{message: message}
}
