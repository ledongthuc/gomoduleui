package flows

import (
	"fmt"

	"github.com/ledongthuc/gomoduleui/models"
	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"golang.org/x/mod/modfile"
)

func ModuleReplaceOptions(previousMenuItem models.MenuItem) (models.MenuItem, error) {
	if previousMenuItem.OriginalData == nil {
		return models.MenuItem{}, errors.New("Can't load library information")
	}
	requiredModule, castable := previousMenuItem.OriginalData.(modfile.Require)
	if !castable {
		return models.MenuItem{}, errors.New("Can't load library information with correct format")
	}
	menuItems := []models.MenuItem{
		models.MenuItem{
			Label: "← back", Type: models.MenuTypeModule,
			OriginalData: previousMenuItem.OriginalData,
		},
	}

	config, err := models.ParseConfigurationsFromFile(".gomoduleui.yml")
	if err != nil {
		return models.MenuItem{}, errors.Wrap(err, "Can't load configuration")
	}
	moduleName := models.ModuleName(requiredModule.Mod.Path)
	if configModule, exist := config.Module[moduleName]; exist {
		for _, defaultValue := range configModule.DefaultReplaces {
			menuItems = append(menuItems, models.MenuItem{
				Label: string(defaultValue), Type: models.MenuTypeModuleReplaceAction,
				OriginalData: previousMenuItem.OriginalData,
			})
		}
	}

	prompt := promptui.Select{
		Label:     fmt.Sprintf("Replace \"%s\" to:", moduleName),
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
