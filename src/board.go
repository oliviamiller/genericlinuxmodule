// Package orangepi implements a orangepi based board.
package orangepi

import (
	"errors"

	"periph.io/x/host/v3"

	"board-module/genericlinux"

	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

var Model = resource.NewModel("viam", "board", "orangepi")

const model = "viam:board:orangepi"

func init() {
	if _, err := host.Init(); err != nil {
		logging.Global().Debugw("error initializing host", "error", err)
	}

	gpioMappings, err := genericlinux.GetGPIOBoardMappings(model, BoardInfoMappings)
	var noBoardErr genericlinux.NoBoardFoundError
	if errors.As(err, &noBoardErr) {
		logging.Global().Debugw("error getting orangepi GPIO board mapping", "error", err)
	}

	genericlinux.RegisterBoard(model, gpioMappings)
}
