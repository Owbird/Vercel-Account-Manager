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
		if err != nil {
			log.Fatalln(err)
		}

		currentAuth, _ := utils.GetVercelDir()

		data, err := os.ReadFile(currentAuth)
		if err != nil {
			log.Fatalln(err)
		}

		var currentConfig map[string]string

		json.Unmarshal(data, &currentConfig)

		currentToken := currentConfig["token"]

		for _, dir := range dirs {
			account := dir.Name()

			authPath := filepath.Join(vamDir, account, "auth.json")

			data, err := os.ReadFile(authPath)
			if err != nil {
				log.Fatalln(err)
			}

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
