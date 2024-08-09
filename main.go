package main

import (
	"github.com/stevenhobs4323/code-mint/internal/base"
	"github.com/stevenhobs4323/code-mint/internal/config"
	"github.com/stevenhobs4323/code-mint/internal/vscode"
)

func main() {
	defer base.LOG_FILE.Close()
	base.CheckApp()
	config.LoadData()
	vscode.Launch()
}
