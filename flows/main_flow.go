package flows

import (
	"github.com/ledongthuc/gomoduleui/commands"
	"github.com/ledongthuc/gomoduleui/models"
)

func MainFlow() error {
	modFile, err := models.GetModFile()
	if err != nil {
		return err
	}

	err = commands.ListModule(modFile)
	if err != nil {
		return err
	}

	return nil
}
