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
	"loganalyzer/logalizer"

	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Print Top 10 status codes",
	Long: `the status code that the server sends back to the client. This information is very valuable, because it reveals whether the request
resulted in a successful response (codes beginning in 2), a redirection (codes beginning in 3), an error caused by the client (codes beginning in 4), 
or an error in the server (codes beginning in 5)`,
	Run: func(cmd *cobra.Command, args []string) {
		logalizer.TopStatusCmd(&filename)
	},
}

func init() {
	topCmd.AddCommand(statusCmd)
}
