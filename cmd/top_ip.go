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

// topIpsCmd represents the topIps command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Print Top 10 IP addresses accessing your web server",
	Long: `Find Top 10 IP Addresses Accessing Your web server,
it helps you to quicly identify abuse`,
	Run: func(cmd *cobra.Command, args []string) {
		topIP(cmd)
	},
}

// Print top 10
func topIP(cmd *cobra.Command) {
	fileName := &filename

	F := logalizer.ReadFile(*fileName)
	ips, _, _, _, _, _, _, _ := logalizer.ParseLog(F)
	ip := logalizer.TopOccurr(ips)
	logalizer.TablePrint("IP", "Count", ip)
}

func init() {
	topCmd.AddCommand(ipCmd)
}
