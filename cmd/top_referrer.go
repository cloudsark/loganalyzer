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
	"github.com/cloudsark/loganalyzer/logalizer"
	"github.com/spf13/cobra"
)

// referrerCmd represents the referrer command
var referrerCmd = &cobra.Command{
	Use:   "referrer",
	Short: "Print top 10 referrers",
	Long: ` referer is an optional HTTP header field that identifies the address of the webpage ... 
	Referer logging is used to allow websites and web servers to identify where people are visiting them from`,
	Run: func(cmd *cobra.Command, args []string) {
		logalizer.TopReferrerCmd(&filename)
	},
}

func init() {
	topCmd.AddCommand(referrerCmd)
}
