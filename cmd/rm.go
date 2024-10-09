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
		utils.HandleFatalError(err)

		for _, dir := range dirs {
			log.Println(dir.Name())
			if account == dir.Name() {
				accountPath := filepath.Join(vamDir, account)

				err := os.RemoveAll(accountPath)
				utils.HandleFatalError(err)

				log.Println("[+] Successfully removed account")

				os.Exit(0)
			}

		}

		log.Fatalln("No account found")
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
