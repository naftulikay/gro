package cmd

import (
  "fmt"
  "os"
  "path/filepath"

  "github.com/spf13/cobra"
)

var (
  FileType string
)

func init() {
  rootCmd.AddCommand(findCmd)

  findCmd.Flags().StringVarP(&FileType, "type", "t", "", "Types of files to match; f = file, d = directory")
}

var findCmd = &cobra.Command{
  Use: "find [path...]",
  Short: "is this even used anywhere?",
  Long: "fucking find shit",
  Run: func(cmd *cobra.Command, args []string) {
    if len(args) > 0 {
      for _, d := range args {
        walk(d)
      }
    } else {
      walk(".")
    }
  },
}

func walk(root string) {
  f, err := os.Open(root)

  if err != nil {
    fmt.Println("error")
    os.Exit(1)
  }

  fmt.Println("starting directory:", f.Name())

  rlistdir(f)
}

func rlistdir(d *os.File) {
  fmt.Println(d.Name() + "/")

  fileinfos, err := d.Readdir(0)

  if err != nil {
    fmt.Println("unable to list directory ", d.Name())
    os.Exit(1)
  }

  for _, info := range fileinfos {
    path := filepath.Join(d.Name(), info.Name())

    if info.IsDir() {
      f, err := os.Open(path)
      if err != nil {
        fmt.Println("unable to open", path)
        os.Exit(1)
      }
      rlistdir(f)
    } else {
      fmt.Println(path)
    }
  }
}
