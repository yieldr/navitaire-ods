// Copyright © 2018 Yieldr

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yieldr/navitaire-ods/pkg/version"
)

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		if short, _ := cmd.Flags().GetBool("short"); short {
			fmt.Print(version.Version)
			return
		}
		fmt.Print(version.LongVersion())
	},
}

func init() {
	cmdRoot.AddCommand(cmdVersion)
	cmdVersion.Flags().BoolP("short", "s", false, "Print a short version")
}
