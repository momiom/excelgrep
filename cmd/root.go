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
	"fmt"
	"os"

	"github.com/momiom/excelgrep/logger"
	"github.com/spf13/cobra"
)

var useVerboseLogger bool
var maxGoroutine int

var Version = "unknown"
var Revision = "unknown"
var showVersion = false

// rootCmd represents the base command when called without any subcommands
func NewCmdRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "eg",
		Short: "excelgrep (eg) recursively searches your current directory and grep the xlsx files.",
		Long: `
USAGE:
		eg [OPTIONS] PATTERN [PATH ...]
		command | eg [OPTIONS] PATTERN

ARGS:
	<PATTERN>	A glob pattern used for searching.
	<PATH>	A file or directory to search.
`,
		Args: cobra.MinimumNArgs(1),
		Run:  runRootCmd,
	}

	// 初期処理
	cobra.OnInitialize(
		func() {
			// ロガーを設定
			if useVerboseLogger {
				logger.SetLogger(logger.Verbose)
			}

			// バージョン表示
			if showVersion {
				fmt.Printf("version: %s\nrevision: %s\n", Version, Revision)
				os.Exit(0)
			}
		},
	)

	// ロガー設定
	rootCmd.PersistentFlags().BoolVar(&useVerboseLogger, "verbose", false, "Enable verbose log")

	// grep の並列処理数
	rootCmd.PersistentFlags().IntVarP(&maxGoroutine, "max-procs", "p", 10, "Run up to specified processes at a time")

	// バージョン表示フラグ
	rootCmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "Show version")

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd := NewCmdRoot()
	rootCmd.SetOutput(os.Stdout)

	if err := rootCmd.Execute(); err != nil {
		rootCmd.SetOutput(os.Stderr)
		rootCmd.Println(err)
		os.Exit(1)
	}

	logger.Debugln("DONE")
}
