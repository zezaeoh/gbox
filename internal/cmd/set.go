package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zezaeoh/gbox/internal/logger"
	"github.com/zezaeoh/gbox/internal/storage"
)

func init() {
	rootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:     "set name data",
	Short:   "Set data to storage",
	Aliases: []string{"s"},
	Args:    cobra.ExactValidArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.Logger()

		stg, err := storage.GetStorage()
		if err != nil {
			log.Errorf("Fail to get storage: %s", err)
			return
		}

		name := args[0]
		data := args[1]

		if err := stg.Set(name, data); err != nil {
			log.Errorf("Fail to set data: %s", err)
			return
		}
		log.Infof("Set: %s", name)
	},
}
