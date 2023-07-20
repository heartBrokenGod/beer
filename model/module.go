package model

type ModuleMeta struct {
	Dir *string // location in the file system
}

type Module struct {
	Name       *string
	GoVersion  *string
	Require    []Dependency
	ModuleMeta *ModuleMeta
}

type Dependency struct {
	ModuleName *string
	Version    *string
}

// =========================================================

type CreateModule struct {
	Dir *string
}

type DeleteModule struct {
	ModuleName *string
}
