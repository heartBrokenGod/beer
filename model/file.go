package model

type FileMeta struct {
	Name         *string
	AbsolutePath *string
	Extension    *string
}

type File struct {
	Directives []string
	Imports    []string
	Constants  []Constant
	Variables  []Variable
	Functions  []Function
	Structs    []Struct
}
