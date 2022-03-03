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
	"github.com/fatih/color"
	"os"

	"github.com/clivewalkden/m2-db-sync/common/validateenvironments"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var source string
var destination string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "m2-db-sync",
	Short: "Helps synchronise Magento 2 databases between servers",
	Long: `A application to ease the way you synchronise your databases
between individual Magento 2 servers.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		error := color.New(color.Bold, color.BgRed, color.FgWhite).PrintlnFunc()
		notice := color.New(color.Bold, color.FgBlue).PrintlnFunc()

		// Check if the config has been loaded
		if viper.ConfigFileUsed() == "" {
			error(" Config file required ")
			os.Exit(0)
		}
		notice("Run the synchronise command here!")
		validateenvironments.main(source, destination)
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

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.MarkPersistentFlagRequired("config")
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

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// Output config file location and all loaded settings
		//fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		//fmt.Fprintln(os.Stdout, viper.AllSettings())
	}
}
