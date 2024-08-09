package vscode

import (
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/stevenhobs4323/code-mint/internal/base"
	"github.com/stevenhobs4323/code-mint/internal/config"
)

func Launch() {
	cmd := exec.Command(path.Join(base.CODE_HOME, "MainApp", "Code.exe"))
	env_path := "PATH=" + config.GetEnvPATHString()
	env_others := config.GetEnvOtrString()
	cmd.Env = append(os.Environ(), env_path)
	cmd.Env = append(cmd.Env, env_others...)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
