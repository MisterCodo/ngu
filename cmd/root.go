package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "ngu",
		Short: "A NGU Industries beacons optimizer",
		Long:  "Provides optimized beacons placement of NGU Industries maps.",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
