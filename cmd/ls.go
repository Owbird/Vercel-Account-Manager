package cmd

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/owbird/vercel-account-manager/utils"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Show added accounts",
	Long:  `List current added accounts with current status`,
	Run: func(cmd *cobra.Command, args []string) {
		vamDir := utils.GetVamDir()

		dirs, err := os.ReadDir(vamDir)
		utils.HandleFatalError(err)

		currentAuth, _ := utils.GetVercelDir()

		data, err := os.ReadFile(currentAuth)
		utils.HandleFatalError(err)

		var currentConfig map[string]string

		json.Unmarshal(data, &currentConfig)

		currentToken := currentConfig["token"]

		for _, dir := range dirs {
			account := dir.Name()

			authPath := filepath.Join(vamDir, account, "auth.json")

			data, err := os.ReadFile(authPath)
			utils.HandleFatalError(err)

			var authConfig map[string]string

			json.Unmarshal(data, &authConfig)

			thisToken := (authConfig["token"])

			if currentToken == thisToken {
				log.Printf("[+] %v (Active)", account)
			} else {
				log.Printf("[+] %v ", account)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
