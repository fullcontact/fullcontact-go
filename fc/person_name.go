package fullcontact

type PersonNameOptions func(name *PersonName)
type PersonName struct {
	Given  string `json:"given,omitempty"`
	Family string `json:"family,omitempty"`
	Full   string `json:"full,omitempty"`
}

func NewPersonName(options ...PersonNameOptions) *PersonName {
	pn := &PersonName{}

	for _, opts := range options {
		opts(pn)
	}
	return pn
}

func WithFull(full string) PersonNameOptions {
	return func(name *PersonName) {
		name.Full = full
	}
}

func WithGiven(given string) PersonNameOptions {
	return func(name *PersonName) {
		name.Given = given
	}
}

func WithFamily(family string) PersonNameOptions {
	return func(name *PersonName) {
		name.Family = family
	}
}
