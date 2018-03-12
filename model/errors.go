package model

type Error interface {
	error

	// Type is the kind of error being returned.
	//
	// It will be unique across all error scenario's
	// and can be used by a client to form their
	// own error messages or make decisions about
	// whether an error needs to be displayed.
	Type() string
}

type ErrJoinedOwnMatch struct{}

func (ErrJoinedOwnMatch) Type() string {
	return "JoinedOwnMatch"
}

func (ErrJoinedOwnMatch) Error() string {
	return "A user can not join a match that they created."
}

type ErrMatchNotFound struct{}

func (ErrMatchNotFound) Type() string {
	return "MatchNotFound"
}

func (ErrMatchNotFound) Error() string {
	return "A match with the given id could not be found."
}

type ErrMatchAlreadyFull struct{}

func (ErrMatchAlreadyFull) Type() string {
	return "MatchAlreadyFull"
}

func (ErrMatchAlreadyFull) Error() string {
	return "The match is already full."
}
