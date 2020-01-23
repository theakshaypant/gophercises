package cmd

import (
  "github.com/spf13/cobra"
)


var cfgFile string


// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
  Use:   "./clitask",
  Short: "./clitask is a CLI for managing your TODOs.",

  // Uncomment the following line if your bare application
  // has an action associated with it:
  //	Run: func(cmd *cobra.Command, args []string) { },
}
