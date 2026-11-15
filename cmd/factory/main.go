package main

import (
 "fmt"
 "os"
 "github.com/atstaeff/ai-factory-cli/internal/factory"
)

func main() {
 if err := factory.Execute(os.Args[1:], os.Stdout, os.Stderr); err != nil {
  fmt.Fprintln(os.Stderr, "factory:", err)
  os.Exit(1)
 }
}
