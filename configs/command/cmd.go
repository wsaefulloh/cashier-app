package command

import (
	"github.com/spf13/cobra"
	"github.com/wsaefulloh/go-solid-principle/configs/db"
)

var initCommand = &cobra.Command{
	Use:   "gorest",
	Short: "Simple restapi with go",
}

func init() {
	initCommand.AddCommand(serveCmd)
	initCommand.AddCommand(db.MigrateUp)
	initCommand.AddCommand(db.MigrateDown)
}

func Run(args []string) error {
	initCommand.SetArgs(args)
	return initCommand.Execute()
}
