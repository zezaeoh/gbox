package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zezaeoh/gbox/internal/logger"
	"github.com/zezaeoh/gbox/internal/storage"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:     "get name",
	Short:   "Get data from storage",
	Aliases: []string{"g"},
	Args:    cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.Logger()

		stg, err := storage.GetStorage()
		if err != nil {
			log.Errorf("Fail to get storage: %s", err)
			return
		}

		name := args[0]

		data, err := stg.Get(name)
		if err != nil {
			log.Errorf("Fail to get data: %s", err)
			return
		}

		fmt.Println(data)
	},
}
