package models

import (
	"fmt"

	"golang.org/x/mod/modfile"
)

const (
	MenuItemLabel       = "%s | %s"
	MenuItemDescription = `  Module name: %s
  Version: %s
  Replacing: %t
  Indirect: %t`
)

type MenuItem struct {
	Label        string
	Description  string
	Type         MenuType
	OriginalData interface{}
}

func CreateMenuItemsFromModFile(modfile *modfile.File) []MenuItem {
	if modfile == nil {
		return []MenuItem{}
	}
	menuItems := make([]MenuItem, 0, len(modfile.Require))
	for _, require := range modfile.Require {
		if require == nil {
			continue
		}

		// TODO: add indirect
		menuItems = append(menuItems, CreateMenuItemFromRequire(*require, false))
	}
	return menuItems
}

func CreateMenuItemFromRequire(require modfile.Require, replaced bool) MenuItem {
	path := require.Mod.Path
	version := require.Mod.Version
	indirect := require.Indirect
	var label, description string

	if len(path) == 0 && len(version) == 0 {
		label = "(Nothing)"
	} else {
		label = fmt.Sprintf(MenuItemLabel, path, version)
	}

	if len(path) == 0 {
		path = "(Nothing)"
	}
	if len(version) == 0 {
		version = "(Nothing)"
	}
	description = fmt.Sprintf(MenuItemDescription, path, version, replaced, indirect)

	return MenuItem{
		Label:       label,
		Description: description,
		Type:        MenuTypeModule,
	}
}
