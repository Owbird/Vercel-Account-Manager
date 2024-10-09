package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/muesli/go-app-paths"
)

func HandleFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CopyFile(src string, dst string) {
	data, err := os.ReadFile(src)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(dst, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func GetVamDir() string {
	vamScope := gap.NewScope(gap.User, "vam")
	vamDirs, err := vamScope.DataDirs()
	if err != nil {
		log.Fatalln(err)
	}
	vamDir := vamDirs[0]

	return vamDir
}

func GetVercelDir() (authPath string, configPath string) {
	vercelScope := gap.NewScope(gap.User, "com.vercel.cli")
	vercelDirs, err := vercelScope.DataDirs()
	if err != nil {
		log.Fatalln(err)
	}
	vercelDir := vercelDirs[0]

	authPath = filepath.Join(vercelDir, "auth.json")
	configPath = filepath.Join(vercelDir, "config.json")

	return authPath, configPath
}
