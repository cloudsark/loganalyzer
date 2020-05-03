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

// requestCmd represents the request command
var requestCmd = &cobra.Command{
	Use:   "request",
	Short: "Print Top 10 requests",
	Long: `The request line contains a great deal of useful information such as the method used by the client
and the resource requested by the client`,
	Run: func(cmd *cobra.Command, args []string) {
		topRequest(cmd)
	},
}

func topRequest(cmd *cobra.Command) {
	fileName := &filename
	var reqList []string
	F := logalizer.ReadFile(*fileName)
	_, _, _, _, _, _, _, requests := logalizer.ParseLog(F)
	for _, request := range requests {
		reqList = append(reqList, request)
	}
	result := logalizer.TopOccurr(reqList)
	logalizer.TablePrint("Method", "Count", result)
}

func init() {
	topCmd.AddCommand(requestCmd)
}
