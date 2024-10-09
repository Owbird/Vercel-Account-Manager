package cmd

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/owbird/vercel-account-manager/utils"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Add a new account",
	Long: `You may also provide your email address, or Team slug (for SAML Single Sign-On), as an argument.
When no arguments are provided, one of the flags are expected to be set.`,
	Run: func(cmd *cobra.Command, args []string) {
		parsedArgs := []string{"login"}

		if len(args) == 1 {
			parsedArgs = append(parsedArgs, args[0])
		}

		if isGithub, err := cmd.Flags().GetBool("github"); err == nil && isGithub {
			parsedArgs = append(parsedArgs, "--github")
		}

		if isGitlab, err := cmd.Flags().GetBool("gitlab"); err == nil && isGitlab {
			parsedArgs = append(parsedArgs, "--gitlab")
		}

		if isBitbucket, err := cmd.Flags().GetBool("bitbucket"); err == nil && isBitbucket {
			parsedArgs = append(parsedArgs, "--bitbucket")
		}

		if len(parsedArgs) == 1 {
			cmd.Help()
			os.Exit(0)
		}

		vCmd := exec.Command("vercel", parsedArgs...)

		vCmd.Stdout = os.Stdout
		vCmd.Stderr = os.Stderr

		vCmd.Run()
		vCmd.Wait()

		vWhoAmI, err := exec.Command("vercel", "whoami").Output()
		utils.HandleFatalError(err)

		currentUser := string(vWhoAmI)

		vamDir := utils.GetVamDir()

		accountDir := filepath.Join(vamDir, currentUser)

		os.MkdirAll(accountDir, os.ModePerm)

		authPath, configPath := utils.GetVercelDir()

		utils.CopyFile(configPath, filepath.Join(accountDir, "config.json"))
		utils.CopyFile(authPath, filepath.Join(accountDir, "auth.json"))

		log.Printf("[+] Account %s added", currentUser)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().Bool("github", false, "Using the vercel login command with the --github option.")
	createCmd.Flags().Bool("gitlab", false, "Using the vercel login command with the --gitlab option.")
	createCmd.Flags().Bool("bitbucket", false, "Using the vercel login command with the --bitbucket option.")
}
