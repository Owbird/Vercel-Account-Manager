/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/owbird/vercel-account-manager/utils"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a saved account",
	Long:  `Permanently delete local account data`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		account := args[0]

		vamDir := utils.GetVamDir()

		dirs, err := os.ReadDir(vamDir)
		if err != nil {
			log.Fatalln(err)
		}

		for _, dir := range dirs {
			log.Println(dir.Name())
			if account == dir.Name() {
				accountPath := filepath.Join(vamDir, account)

				err := os.RemoveAll(accountPath)
				if err != nil {
					log.Fatalln(err)
				}

				log.Println("[+] Successfully removed account")

				os.Exit(0)
			}

		}

		log.Fatalln("No account found")
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
