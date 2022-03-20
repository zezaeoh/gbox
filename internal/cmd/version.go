package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zezaeoh/gbox/internal/common"
	"github.com/zezaeoh/gbox/internal/logger"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version of gbox cli",
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.Logger()
		log.Infof("GBox Version = %s", common.GetVersion())
	},
}
