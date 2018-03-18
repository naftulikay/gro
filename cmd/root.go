package cmd

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "gro",
  Short: "A crommand line utilidang in grolang.",
  Long: "A crommand line utilidang in grolang.",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Hello world.")
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
