package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zezaeoh/gbox/internal/cmd/storage"
)

func init() {
	rootCmd.AddCommand(storage.Cmd)
}

var rootCmd = &cobra.Command{
	Use:   "gbox",
	Short: "Gbox is simple storage cli using github as storage 📦",
	Args:  cobra.MinimumNArgs(1),
}

func Execute() error {
	return rootCmd.Execute()
}
