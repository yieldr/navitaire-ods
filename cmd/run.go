// Copyright © 2018 Yieldr

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yieldr/navitaire-ods/pkg/navitaire/ods"
	"github.com/yieldr/navitaire-ods/pkg/yieldr"
)

var argRun struct {
}

var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run the uploader",
	Long: `Yieldr - Navitaire ODS Flight Uploader

This program queries the Navitaire ODS database for flight revenue and load
factor per inventory leg and subsequently uploads the result to Yieldr using
either the Yieldr API or SFTP.`,
	Run: run,
}

func run(cmd *cobra.Command, args []string) {

	o, err := ods.New(&ods.ODSConfig{
		Driver:   viper.GetString("db"),
		Addr:     viper.GetString("db-addr"),
		User:     viper.GetString("db-user"),
		Password: viper.GetString("db-pass"),
		Database: viper.GetString("db-name"),
	})
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed connecting to Navitaire ODS. %s", err)
		os.Exit(1)
	}

	query := viper.GetString("db-query")
	queryArgs := viper.GetStringSlice("db-query-args")

	var q []byte

	switch query {
	case "-":
		q, err = ioutil.ReadAll(os.Stdin)
	case "":
		q = ods.DefaultQuery()
	default:
		q, err = ioutil.ReadFile(query)
	}

	if err != nil {
		fmt.Fprintf(os.Stdout, "failed reading query. %s", err)
		os.Exit(1)
	}

	flights, err := o.Query(string(q), queryArgs...)
	if err != nil {
		fmt.Fprintf(os.Stdout, "failed querying for flights. %s", err)
		os.Exit(1)
	}

	if viper.GetBool("api") {

		var buf bytes.Buffer
		e := json.NewEncoder(&buf)
		for _, f := range flights {
			e.Encode(f)
		}

		yldr := yieldr.New(&yieldr.YieldrConfig{
			Addr:         viper.GetString("api-addr"),
			ClientID:     viper.GetString("api-client-id"),
			ClientSecret: viper.GetString("api-client-secret"),
		})

		projectID := viper.GetInt("api-project-id")
		if projectID == -1 {
			fmt.Fprintf(os.Stdout, "The --api-project-id flag is required\n")
			os.Exit(1)
		}

		err = yldr.Upload(projectID, buf.Bytes())
		if err != nil {
			fmt.Fprintf(os.Stdout, "failed uploading flights to Yieldr. %s", err)
			os.Exit(1)
		}
	}

	if viper.GetBool("sftp") {
		fmt.Fprintf(os.Stderr, "SFTP upload is not supported yet!")
		os.Exit(1)
	}
}

func init() {
	cmdRoot.AddCommand(cmdRun)

	cmdRun.Flags().Bool("api", true, "Use the api to upload flights")
	cmdRun.Flags().String("api-addr", "myaccount.yieldr.com", "API server address")
	cmdRun.Flags().String("api-client-id", "", "API client id")
	cmdRun.Flags().String("api-client-secret", "", "API client secret")
	cmdRun.Flags().Int("api-project-id", -1, "API project id")
	viper.BindPFlag("api", cmdRun.Flags().Lookup("api"))
	viper.BindPFlag("api-addr", cmdRun.Flags().Lookup("api-addr"))
	viper.BindPFlag("api-client-id", cmdRun.Flags().Lookup("api-client-id"))
	viper.BindPFlag("api-client-secret", cmdRun.Flags().Lookup("api-client-secret"))
	viper.BindPFlag("api-project-id", cmdRun.Flags().Lookup("api-project-id"))

	cmdRun.Flags().Bool("sftp", false, "Use sftp to upload flights")
	cmdRun.Flags().String("sftp-addr", "localhost:22", "SFTP server address")
	cmdRun.Flags().String("sftp-user", "", "SFTP user")
	cmdRun.Flags().String("sftp-key-file", "", "SFTP key file in PEM format")
	viper.BindPFlag("sftp", cmdRun.Flags().Lookup("sftp"))
	viper.BindPFlag("sftp-addr", cmdRun.Flags().Lookup("sftp-addr"))
	viper.BindPFlag("sftp-user", cmdRun.Flags().Lookup("sftp-user"))
	viper.BindPFlag("sftp-key-file", cmdRun.Flags().Lookup("sftp-key-file"))

	cmdRun.Flags().String("db", "sqlserver", "Database driver")
	cmdRun.Flags().String("db-addr", "localhost:1234", "Database server address")
	cmdRun.Flags().String("db-user", "", "Database user")
	cmdRun.Flags().String("db-pass", "", "Database password")
	cmdRun.Flags().String("db-name", "", "Database name")
	cmdRun.Flags().String("db-query", "", "SQL file to run against the database. Use a dash (-) for stdin or leave empty for the default query")
	cmdRun.Flags().StringSlice("db-query-args", []string{}, "Arguments to the sql query")
	viper.BindPFlag("db", cmdRun.Flags().Lookup("db"))
	viper.BindPFlag("db-addr", cmdRun.Flags().Lookup("db-addr"))
	viper.BindPFlag("db-user", cmdRun.Flags().Lookup("db-user"))
	viper.BindPFlag("db-pass", cmdRun.Flags().Lookup("db-pass"))
	viper.BindPFlag("db-name", cmdRun.Flags().Lookup("db-name"))
	viper.BindPFlag("db-name", cmdRun.Flags().Lookup("db-name"))
	viper.BindPFlag("db-query", cmdRun.Flags().Lookup("db-query"))
	viper.BindPFlag("db-query-args", cmdRun.Flags().Lookup("db-query-args"))
}
