// Copyright © 2018 The Kubicorn Authors
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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/kubicorn/controller/loop"
	"github.com/kubicorn/kubicorn/pkg/logger"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubicorn-controller",
	Short: "The Kubicorn machine controller",
	Long: `Run the Kubicorn controller to reconcile your infrastructure like the beautiful person you are.`,
	Run: func(cmd *cobra.Command, args []string) {



		// TODO @kris-nova
		// We need to hack in here with the real control loop (once I write it...)
		loop.RunService()
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

	rootCmd.PersistentFlags().IntVarP(&logger.Level, "verbose", "v", 4, "Log level")

}