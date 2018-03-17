// Copyright © 2018 Yieldr

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"
)

var cmdDoc = &cobra.Command{
	Use:   "doc",
	Short: "Generate cli documentation",
	Run:   runDoc,
}

func runDoc(cmd *cobra.Command, args []string) {
	dir, _ := cmd.Flags().GetString("dir")
	format, _ := cmd.Flags().GetString("format")

	var err error

	switch format {
	case "markdown":
		err = doc.GenMarkdownTree(cmdRoot, dir)
	case "rest":
		err = doc.GenReSTTree(cmdRoot, dir)
	case "man":
		err = doc.GenManTree(cmdRoot, &doc.GenManHeader{}, dir)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cmdRoot.AddCommand(cmdDoc)
	cmdDoc.Flags().String("dir", "doc", "Directory in which to generate documentation")
	cmdDoc.Flags().String("format", "markdown", "Format in which to generate documentation")
	viper.BindPFlag("dir", cmdRun.Flags().Lookup("dir"))
	viper.BindPFlag("format", cmdRun.Flags().Lookup("format"))
}
