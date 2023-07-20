package model

type Function struct {
	Name         *string
	InputParams  []InputParam
	ReturnParams []ReturnParam
}

type InputParam struct {
	Name *string
	Type *string
}

type ReturnParam struct {
	Name *string
	Type *string
}
