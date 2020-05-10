package flows

import "github.com/ledongthuc/gomoduleui/models"

func StartFlow() error {
	menuItem := models.GetStartMenuItem()
	var err error
	for {
		menuItem, err = RunFlowFromSelection(menuItem)
		if err != nil {
			return err
		}
		if menuItem.Type == models.MenuTypeExit {
			return nil
		}
	}
}

func RunFlowFromSelection(menuItem models.MenuItem) (models.MenuItem, error) {
	switch menuItem.Type {
	case models.MenuTypeStart:
		return ListOptions()
	case models.MenuTypeModule:
		return ModuleManagementOptions(menuItem)
	case models.MenuTypeModuleReplaceOptions:
		return ModuleReplaceOptions(menuItem)
		// case models.MenuTypeModuleUpdateOptions:
		// 	return ModuleUpdateOptions(menuItem)
	}
	return models.GetStartMenuItem(), nil
}
