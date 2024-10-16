package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/owbird/vercel-account-manager/utils"
	"github.com/spf13/cobra"
)

// checkoutCmd represents the checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout",
	Short: "Checkout a saved account",
	Long:  `Switch the current account to one the saved accounts`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		account := args[0]

		vamDir := utils.GetVamDir()

		accountPath := filepath.Join(vamDir, account)

		if _, err := os.Stat(accountPath); os.IsNotExist(err) {
			log.Fatalf("Account %s does not exist", account)
		}

		vercelAuthPath, vercelConfigPath := utils.GetVercelDir()

		utils.CopyFile(filepath.Join(accountPath, "auth.json"), vercelAuthPath)
		utils.CopyFile(filepath.Join(accountPath, "config.json"), vercelConfigPath)

		log.Printf("Successfully switched to %s", account)
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
}
