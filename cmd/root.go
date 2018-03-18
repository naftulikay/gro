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
    f, err := os.Open(".ansible")
    if err != nil {
      fmt.Println("error")
      os.Exit(1)
    }

    rlistdir(f)
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func rlistdir(d *os.File) {
  fmt.Println(d.Name() + "/")

  fileinfos, err := d.Readdir(0)

  if err != nil {
    fmt.Println("nope")
    os.Exit(1)
  }

  for _, info := range fileinfos {
    if info.IsDir() {
      f, err := os.Open(d.Name() + info.Name())
      if err != nil {
        fmt.Println("unable to open")
        os.Exit(1)
      }
      rlistdir(f)
    } else {
      fmt.Println(d.Name() + "/" + info.Name())
    }
  }
}
