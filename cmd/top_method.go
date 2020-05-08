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
	"github.com/mohtork/loganalyzer/logalizer"

	"github.com/spf13/cobra"
)

// methodCmd represents the method command
var methodCmd = &cobra.Command{
	Use:   "method",
	Short: "Print Top 10 HTTP methods used",
	Long:  `HTTP defines a set of request methods to indicate the desired action to be performed for a given resource`,
	Run: func(cmd *cobra.Command, args []string) {
		logalizer.TopMethodCmd(&filename)
	},
}

func init() {
	topCmd.AddCommand(methodCmd)
}
