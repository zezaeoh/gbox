package storage

import (
	"github.com/spf13/cobra"
)

// export root cmd
var Cmd = rootCmd

var rootCmd = &cobra.Command{
	Use:     "storage",
	Aliases: []string{"s", "stg"},
	Short:   "Get, Set storage configuration of gbox",
	Args:    cobra.MinimumNArgs(1),
}
