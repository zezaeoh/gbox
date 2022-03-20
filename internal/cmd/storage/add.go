package storage

import (
	"github.com/spf13/cobra"
	"github.com/zezaeoh/gbox/internal/logger"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add storage config",
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.Logger()
		log.Infof("add method!")
	},
}
