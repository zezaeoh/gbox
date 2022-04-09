package storage

import (
	"github.com/spf13/cobra"
	"github.com/zezaeoh/gbox/internal/logger"
	"github.com/zezaeoh/gbox/internal/storage"
)

func init() {
	rootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:   "set name",
	Short: "Set current storage",
	Args:  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.Logger()

		cfg, err := storage.GetConfig()
		if err != nil {
			log.Errorf("Fail to get config: %s", err)
			return
		}

		name := args[0]
		nameExist := false
		for _, s := range cfg.Storages {
			if s.Name == name {
				nameExist = true
				break
			}
		}
		if !nameExist {
			log.Errorf("There is no named storage in config: %s", name)
			return
		}

		cfg.CurrentStorage = name
		if err := cfg.Save(); err != nil {
			log.Errorf("Error while saving config: %s", err)
			return
		}
		log.Infof("Storage Configured: %s", name)
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		cfg, err := storage.GetConfig()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		return cfg.GetMatchedStorage(toComplete), cobra.ShellCompDirectiveNoFileComp
	},
}
