package errs

// CustomError is a custom struct for handling error messages.
type CustomError struct {
	Text   string `json:"text"`
	Status int    `json:"status"`
}

func (e CustomError) Error() string {
	return e.Text
}

func New(text string, status int) CustomError {
	return CustomError{text, status}
}
