// cmd/git-pitch/main.go
package main

import "github.com/guilhermehbueno/git-pitch/cmd" // adjust module path

var (
  version = "dev"
  commit  = "none"
  date    = "unknown"
)

func main() { cmd.Execute(version, commit, date) }