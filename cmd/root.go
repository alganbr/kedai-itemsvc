package cmd

import (
	"github.com/alganbr/kedai-itemsvc/internal/server"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var RootCmd = &cobra.Command{
	Use:   "itemsvc",
	Short: "Item service API",
	Long:  "Item service API",
	RunE: func(cmd *cobra.Command, args []string) error {
		fx.New(server.Module).Run()
		return nil
	},
}
