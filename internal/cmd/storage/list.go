package storage

import (
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
	Short:   "List storage configs",
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.Logger()

		cfg, err := storage.GetConfig()
		if err != nil {
			log.Errorf("Fail to get config: %s", err)
			return
		}

		if len(cfg.Storages) == 0 {
			log.Warn("empty")
			return
		}

		for _, s := range cfg.Storages {
			if s.Name == cfg.CurrentStorage {
				log.Info("✓ " + s.Name)
			} else {
				log.Info(s.Name)
			}
		}
	},
}
