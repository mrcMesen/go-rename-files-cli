/*
Copyright © 2023 MARCO MESEN  marcokse@hotmail.es
*/
package cmd

import (
	"fmt"
	"github/mrcmesen/go-rename-files-cli/utils"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-rename-files-cli",
	Short: "This is a tool to rename files",
	Long: `
This is a CLI Tool to rename file from camelCase or PascalCase to kebab-case. 

- Give the path of the folder where the files are located
- We will list the files and folders in that path
- We will ask you if you want to rename the files

`,
	Run: func(cmd *cobra.Command, args []string) {
		path := utils.GetPathFromUserPrompt()
		filesToRename := utils.ListCamelCasePaths(path)
		for len(filesToRename) == 0 {
			fmt.Println("No files to rename")
			path = utils.GetPathFromUserPrompt()
			filesToRename = utils.ListCamelCasePaths(path)
		}
		filesRenamed := utils.ListNewNames(filesToRename)
		utils.PrintList(filesToRename, filesRenamed)
		if utils.GetConfirmationFromUserPrompt() {
			// rename all the files
			for _, oldPath := range filesToRename {
				newPath := utils.GetSnakeCase(oldPath)
				err := utils.RenameFile(oldPath, newPath)
				if err != nil {
					panic(err)
				}
			}
			// print the number of files renamed
			fmt.Printf("Renamed %d files\n", len(filesToRename))
			// celebrate with a message and emoji
			fmt.Println("🎉 All files renamed successfully!")
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-rename-files-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
