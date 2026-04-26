package cmd

import (
	"github.com/c1r5/url-shortener/src/app"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "url-shortener",
	Short: "A simple URL shortener service",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.Run()
	},
}

func Execute() error {
	return rootCmd.Execute()
}
