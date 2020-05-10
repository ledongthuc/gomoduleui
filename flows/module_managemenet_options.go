package flows

import (
	"github.com/ledongthuc/gomoduleui/models"
	"github.com/manifoldco/promptui"
)

func ModuleManagementOptions(previousMenuItem models.MenuItem) (models.MenuItem, error) {
	menuItems := []models.MenuItem{
		models.MenuItem{
			Label: "← back", Type: models.MenuTypeStart,
		},
		models.MenuItem{
			Label: "Replace module", Type: models.MenuTypeModuleReplaceOptions,
			OriginalData: previousMenuItem.OriginalData,
		},
		models.MenuItem{
			Label: "Update version", Type: models.MenuTypeModuleUpdateOptions,
			OriginalData: previousMenuItem.OriginalData,
		},
	}

	prompt := promptui.Select{
		Label:     "What do you want do to?",
		Items:     menuItems,
		Templates: models.GetDefaultMenuTemplate(),
		Size:      7,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return models.MenuItem{}, err
	}
	return menuItems[i], nil
}
