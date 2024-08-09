package config

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/stevenhobs4323/code-mint/internal/base"
)

func GetEnvPATHString() string {
	originPath := os.Getenv("PATH")
	for binDir, enable := range config.AppendPath {
		// 判断 binDir 是否为绝对路径
		if !filepath.IsAbs(binDir) {
			binDir = path.Join(base.CODE_HOME, binDir)
		}
		if enable == 1 {
			originPath = binDir + ";" + originPath
		}
	}
	return strings.Replace(originPath, "/", `\`, -1)
}
func GetEnvOtrString() (envs []string) {
	for key, value := range config.OtherEnv {
		// 判断 value 是否为绝对路径
		if !filepath.IsAbs(value) {
			value = path.Join(base.CODE_HOME, value)
		}
		envs = append(envs, key+"="+strings.ReplaceAll(value, "/", `\`))
	}
	return
}
