// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
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
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var (
	domain string
)

// create new http client
var client = http.Client{
	Transport: &http.Transport{},
}

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check the status of a website",
	Args: func(checkCmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("At least one argument is required")
		}
		domain = args[0]
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check called")
		site := "https://" + domain
		request, err := http.NewRequest("HEAD", site, nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't make new request: %v\n", err)
			os.Exit(1)
		}
		response, err := client.Do(request)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't send request: %v\n", err)
			os.Exit(1)
		}
		response.Body.Close()
		fmt.Println("Status: ", response.StatusCode)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
