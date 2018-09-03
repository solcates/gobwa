// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/solcates/gobwa/pkg/bwa"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var debug bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gobwa",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// See if we got a target IP/HOST from the CLI

		if debug {
			logrus.SetLevel(logrus.DebugLevel)
		}

		var targetSpa string
		if len(args) != 1 {
			//bc := broadcast.NewUDPBroadcaster(30303, "Discovery: Who is out there?")
			bc := bwa.NewDiscoverer()
			spas, err := bc.Discover()
			if err != nil {
				logrus.Fatal(err)
			}
			switch len(spas) {
			case 0:
				logrus.Fatal("No SPA Found on network.  Check that we are on the same WIFI AP as spa?")
			case 1:
				targetSpa = spas[0]
			default:
				logrus.Warn("Found more than 1 Spa on your network.  Rerun command with one of these IP addresses.")
				for _, spa := range spas {
					logrus.Warn(spa)
				}
			}
		} else {
			//
			targetSpa = args[0]
		}

		// Setup and use the client
		client := bwa.NewBalbowClient(targetSpa, 4257)
		client.Connect()
		client.RequestConfig()
		for {
			select {
			case <-time.After(5 * time.Second):
				//client.RequestControlInfo()
				client.ToggleLight()
				break

			}
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gobwa.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Debug mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".gobwa" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gobwa")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
