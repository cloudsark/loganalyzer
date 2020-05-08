/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var filename string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "loganalyzer",
	Short: "Analyze access logs",
	Version: "0.2",  // hard code that for now
	Long:  `Access log analyzer can be used to analyze web servers , load balancers access logs `,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Something's wrong--perhaps a missing command")
		}
		return fmt.Errorf("invalid command: %s", args[0])
	},

	Run: func(cmd *cobra.Command, args []string) {
		filename, err := cmd.Flags().GetString("file")

		if filename == "" {
			log.Fatalf("failed to open file: %s", err)
		}

		filenameAbsPath, err := filepath.Abs(filename)
		if err != nil {
			log.Fatalf("failed to open file: %s", err)
		}
		filename = filenameAbsPath
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&filename, "file", "f", "", "access log filename")
	rootCmd.MarkPersistentFlagRequired("file")
	rootCmd.MarkFlagFilename("file")
}
