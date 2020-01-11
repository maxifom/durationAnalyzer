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
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// parseDirCmd represents the parseDir command
var parseDirCmd = &cobra.Command{
	Use:   "parseDir",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		filename, err := cmd.Flags().GetString("dir")
		if err != nil {
			return err
		}
		filenameAbsPath, err := filepath.Abs(filename)
		if err != nil {
			return err
		}
		filepath.Walk(filenameAbsPath, func(path string, info os.FileInfo, err error) error {
			if !strings.HasSuffix(path, ".go") || info.IsDir() {
				return nil
			}
			log.Println(path)
			err = analyze(path)
			if err != nil {
				return err
			}
			return nil
		})
		return nil
	},
}

func init() {
	rootCmd.AddCommand(parseDirCmd)
	parseDirCmd.Flags().StringP("dir", "d", "", "directory to analyze")
	parseDirCmd.MarkFlagFilename("dir")
}
