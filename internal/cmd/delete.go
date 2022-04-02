package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zezaeoh/gbox/internal/logger"
	"github.com/zezaeoh/gbox/internal/storage"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:     "delete name",
	Short:   "Delete data from storage",
	Aliases: []string{"d"},
	Args:    cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.Logger()

		stg, err := storage.GetStorage()
		if err != nil {
			log.Errorf("Fail to get storage: %s", err)
			return
		}

		name := args[0]

		if err := stg.Delete(name); err != nil {
			log.Errorf("Fail to delete data: %s", err)
			return
		}
		log.Infof("Delete: %s", name)
	},
}
