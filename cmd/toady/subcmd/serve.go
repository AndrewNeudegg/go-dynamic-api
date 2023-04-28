package serve

import (
	"github.com/spf13/cobra"
)

// Cmd demonstrates how to configure a new subcommand.
func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "serve event routing as described by the given configuration",
		Long:  `serve will orchestrate whatever recipe has been described by your configuration.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}

	// cmd.Flags().StringVarP(&configurationPath, "config", "c", "", "the application configuration, if using stdin specify '-'.")

	return cmd
}
