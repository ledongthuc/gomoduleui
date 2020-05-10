package flows

import (
	"github.com/ledongthuc/gomoduleui/models"
	"github.com/manifoldco/promptui"
)

func ListOptions() (models.MenuItem, error) {
	modFile, err := models.GetModFile()
	if err != nil {
		return models.MenuItem{}, err
	}

	menuItems := models.CreateMenuItemsFromModFile(modFile)
	menuItems = append(menuItems, models.MenuItem{
		Label: "← exit",
		Type:  models.MenuTypeExit,
	})

	prompt := promptui.Select{
		Label:     "Select:",
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
