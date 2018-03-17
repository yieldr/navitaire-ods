// Copyright © 2018 Yieldr

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cmdRoot = &cobra.Command{
	Use:   "navitaire-ods",
	Short: "Root command",
	Long: `Yieldr - Navitaire ODS Flight Uploader

Use this program to query your Navitaire ODS database for flight performance and
upload the result to Yieldr.

See the 'run' sub command for more details.
`,
}

func Execute() {
	if err := cmdRoot.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		viper.AutomaticEnv()
		viper.SetEnvPrefix("YIELDR")
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	})

	cmdRoot.PersistentFlags().BoolP("debug", "d", false, "Enable debug mode to print more verbose logs")
	viper.BindPFlag("debug", cmdRoot.PersistentFlags().Lookup("debug"))
}
