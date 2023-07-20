package model

type Directory struct {
	Name *string
	Path *string
}

// api structs
type CreateDirectory struct {
	Name     *string
	Location *string
}
