package models

type MenuType int

const (
	MenuTypeStart                MenuType = iota
	MenuTypeExit                 MenuType = iota
	MenuTypeModule               MenuType = iota
	MenuTypeModuleReplaceOptions MenuType = iota
	MenuTypeModuleReplaceAction  MenuType = iota
	MenuTypeModuleUpdateOptions  MenuType = iota
	MenuTypeCustom               MenuType = iota
)
