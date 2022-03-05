package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "trim-head-space",
  Short: "trim-head-space trims head spaces",
  Long: "trim-head-space trims head spaces",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("args", args)
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
