// package main is a module with correction-station component
package main

import (
	orangepi "board-module/src"
	"context"

	"go.viam.com/rdk/components/board"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/module"
	"go.viam.com/utils"
)

func main() {
	utils.ContextualMain(mainWithArgs, logging.NewLogger("board"))
}

func mainWithArgs(ctx context.Context, args []string, logger logging.Logger) error {
	board1, err := module.NewModuleFromArgs(ctx, logger)

	if err != nil {
		return err
	}
	board1.AddModelFromRegistry(ctx, board.API, orangepi.Model)

	err = board1.Start(ctx)
	defer board1.Close(ctx)
	if err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}
