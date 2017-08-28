package menu

import (
	"os"

	"github.com/pkg/errors"
	"github.com/vlad-s/hcpxread/helpers"
	"github.com/vlad-s/hcpxread/structs"
)

func ParseChoice(choice int, instances structs.HccapxInstances) error {
	switch true {
	case choice == 0:
		helpers.Logger.Info("Exiting, goodbye")
		os.Exit(0)
	case choice < 0 || choice > len(instances):
		return errors.New("Invalid index")
	default:
		helpers.ClearScreen(true)
		helpers.PrintHccapx(instances[choice-1])
	}
	return nil
}
