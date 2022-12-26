package exception

type DomainError struct {
	message string
	name    string
}

func (e DomainError) Error() string {
	return e.message
}

func NewDomainError(message string, name string) error {
	return &DomainError{message: message, name: name}
}
