package errors

type MissingParamsError struct {
	Code   int
	Params []string
}

func (err MissingParamsError) Error() {

}
