package menu

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/vlad-s/hcpxread/helpers"
	"github.com/vlad-s/hcpxread/structs"
)

func ParseChoice(choice int, instances structs.HccapxInstances) (bool, error) {
	switch true {
	case choice == 0:
		helpers.Logger.Info("Exiting, goodbye")
		os.Exit(0)
	case choice == 99:
		fmt.Printf("export > ")
		fmt.Fscanf(os.Stdin, "%d", &choice)
		if choice < 0 || choice > len(instances) {
			return false, errors.New("Invalid index")
		}

		var fileName string
		fmt.Printf("filename > ")
		fmt.Fscanf(os.Stdin, "%s", &fileName)
		file, err := os.Create(fileName)
		if err != nil {
			return false, errors.Wrap(err, "Error creating the file")
		}
		defer file.Close()
		n, err := file.Write(instances[choice-1].Content)
		if err != nil {
			return false, errors.Wrap(err, "Error writing the file")
		}
		if n != 393 {
			return false, errors.New("Couldn't write the whole content")
		}
		return true, nil
	case choice < 0 || choice > len(instances):
		return false, errors.New("Invalid index")
	default:
		helpers.ClearScreen(true)
		helpers.PrintHccapx(instances[choice-1])
	}
	return false, nil
}
