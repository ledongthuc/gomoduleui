package commands

import (
	"fmt"

	"github.com/ledongthuc/gomoduleui/models"
	"github.com/manifoldco/promptui"
	"golang.org/x/mod/modfile"
)

func ListModule(mod *modfile.File) error {
	actions := models.CreateMenuItemsFromModFile(mod)

	prompt := promptui.Select{
		Label:     "Select action:",
		Items:     actions,
		Templates: models.GetDefaultMenuTemplate(),
		Size:      7,
	}

	i, _, err := prompt.Run()
	if err != nil {
		return err
	}

	fmt.Printf("You choose number: %s\n", actions[i].Label)
	return nil
}
