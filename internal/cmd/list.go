package cmd

import (
	"github.com/shivamMg/ppds/tree"
	"github.com/spf13/cobra"
	"github.com/zezaeoh/gbox/internal/logger"
	"github.com/zezaeoh/gbox/internal/storage"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "l"},
	Short:   "List storage data",
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.Logger()

		stg, err := storage.GetStorage()
		if err != nil {
			log.Errorf("Fail to get storage: %s", err)
			return
		}

		data, err := stg.List()
		if err != nil {
			log.Errorf("Fail to list data: %s", err)
			return
		}

		tree.PrintHrn(data)
	},
}
