package cmd

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/owbird/vercel-account-manager/utils"
	"github.com/spf13/cobra"
)

func copyFile(src string, dst string) {
	data, err := os.ReadFile(src)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(dst, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Add a new account",
	Long:  `Add an existing vercel account to the containers and can be checkout later`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		vCmd := exec.Command("vercel", "login", args[0])

		vCmd.Stdout = os.Stdout
		vCmd.Stderr = os.Stderr

		vCmd.Run()
		vCmd.Wait()

		vamDir := utils.GetVamDir()

		accountDir := filepath.Join(vamDir, args[0])

		os.MkdirAll(accountDir, os.ModePerm)

		authPath, configPath := utils.GetVercelDir()

		copyFile(configPath, filepath.Join(accountDir, "config.json"))
		copyFile(authPath, filepath.Join(accountDir, "auth.json"))

		log.Printf("[+] Account %s added", args[0])
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
