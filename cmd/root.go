/*
Copyright © 2022 Clive Walkden clivewalkden@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"github.com/clivewalkden/m2-db-sync/common"
	"github.com/clivewalkden/m2-db-sync/validation"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string
var source string
var destination string
var full bool
var prefix bool
var wordpress bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "m2-db-sync",
	Short: "Helps synchronise Magento 2 databases between servers",
	Long: `A application to ease the way you synchronise your databases
between individual Magento 2 servers.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// Check if the config has been loaded
		if viper.ConfigFileUsed() == "" {
			common.Error(" Config file required ")
			os.Exit(1)
		}

		errEnvVal := validation.EnvironmentsValidation(source, destination)
		if errEnvVal != nil {
			common.Error(errEnvVal.Error())
			os.Exit(1)
		}

		// Prepare Server values
		config := common.Config{
			Src:       common.ServerSetup(source),
			Dest:      common.ServerSetup(destination),
			PiiDB:     full,
			Prefix:    prefix,
			WordPress: wordpress,
		}

		errConVal := validation.ConfigValidation(config)
		if errConVal != nil {
			common.Error(errConVal.Error())
			os.Exit(1)
		}

		// Test Connection to the server
		//common.Connect(srcServer, "ls -alh")
		common.RemoteBackup(config)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./.m2-db-sync.yaml)")

	rootCmd.PersistentFlags().StringVarP(&source, "source", "s", "", "source environment (production, staging, development)")
	rootCmd.PersistentFlags().StringVarP(&destination, "destination", "d", "", "destination environment (local, development, staging)")
	rootCmd.PersistentFlags().BoolVarP(&full, "full", "f", false, "full database dump (off by default)")
	rootCmd.PersistentFlags().BoolVarP(&prefix, "prefix", "p", true, "prefix invoices, orders, shipments etc (off by default)")
	rootCmd.PersistentFlags().BoolVarP(&wordpress, "wordpress", "w", true, "include WordPress content (off by default)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//rootCmd.MarkPersistentFlagRequired("config")
	rootCmd.MarkPersistentFlagRequired("source")
	rootCmd.MarkPersistentFlagRequired("destination")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in current directory with name ".m2-db-sync" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".m2-db-sync")
	}

	viper.AutomaticEnv() // read in environment variables that match
	//fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	//viper.Debug()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// Output config file location and all loaded settings
		//fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		//fmt.Fprintln(os.Stdout, viper.AllSettings())
	}
}
