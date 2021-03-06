package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var ProblemWithConfigFile bool

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "centralconfig",
	Short: "A simple REST service and UI for application configuration",
	Long: `Centralconfig is a REST based service for managing application configuration.
It's designed to be used with one of many different SQL (or nosql) backends.  You can
use both an API and a web UI to manage configuration information.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/centralconfig.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	viper.AutomaticEnv() // read in environment variables that match

	//	Set our defaults
	viper.SetDefault("server.port", "3000")
	viper.SetDefault("server.bind", "")
	viper.SetDefault("server.allowed-origins", "*")
	viper.SetDefault("datastore.type", "boltdb")
	viper.SetDefault("datastore.database", "config.db")

	viper.SetConfigName("centralconfig") // name of config file (without extension)
	viper.AddConfigPath("$HOME")         // adding home directory as first search path
	viper.AddConfigPath(".")             // also look in the working directory

	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	// If a config file is found, read it in
	// otherwise, make note that there was a problem
	if err := viper.ReadInConfig(); err != nil {
		ProblemWithConfigFile = true
	}
}
