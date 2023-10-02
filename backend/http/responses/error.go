package responses

type Error struct {
	Error string `json:"error"`
}

func newErrorResponse(e string) Error {
	return Error{
		Error: e,
	}
}

func newErrorResponseE(err error) Error {
	return Error{
		Error: err.Error(),
	}
}
