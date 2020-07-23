package fullcontact

import "fmt"

type FullContactError struct {
	message string
}

func (fcError *FullContactError) Error() string {
	return fmt.Sprintf("FullContactError: %s", fcError.message)
}

func NewFullContactError(msg string) *FullContactError {
	return &FullContactError{message: msg}
}

func isPopulated(value string) bool {
	return len(value) > 0
}
