// cmd/root.go
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	ver = "dev"
	com = "none"
	dat = "unknown"
)

func Execute(version, commit, date string) {
	ver, com, dat = version, commit, date
	cobra.CheckErr(rootCmd.Execute())
}

var rootCmd = &cobra.Command{
	Use:   "git-pitch",
	Short: "git-pitch CLI",
}

func init() {
	rootCmd.Version = fmt.Sprintf("%s (%s) %s", ver, com, dat)
}
