package base

import (
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"unicode"

	"github.com/stevenhobs4323/code-mint/internal/utils"
)

var CODE_HOME string
var LOG_FILE *os.File

const configJsonTemplate = `
{
  "AppendPath": {
    "Sdk/bin": 1
  },
  "OtherENV": {
    "XDG_CACHE_HOME": "Cache",
    "XDG_CONFIG_HOME": "Config",
    "TEMP": "Cache/Temp",
    "TMP": "Cache/Temp",
    "USERPROFILE": "UserData",
    "APPDATA": "UserData/Roaming",
    "LOCALAPPDATA": "UserData/Local"
  }
}
`

func init() {
	// 设置应用主文件夹路径
	execPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	CODE_HOME = path.Join(filepath.Dir(execPath), ".mint")
	// CODE_HOME = "absolute_path_to_project"+".mint" //! 测试路径
	// Check CODE_HOME
	if _, err = os.Stat(CODE_HOME); os.IsNotExist(err) {
		err = os.Mkdir(CODE_HOME, 0755)
		if err != nil {
			panic(err)
		}
	}
	// 设置日志文件输出
	logfilePath := path.Join(CODE_HOME, "code-mint.log")
	// 打开日志文件
	LOG_FILE, err := os.OpenFile(logfilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(LOG_FILE)
}

func CheckApp() {
	// Check Cache Folder
	if _, err := os.Stat(path.Join(CODE_HOME, "Cache")); os.IsNotExist(err) {
		err = os.MkdirAll(path.Join(CODE_HOME, "Cache"), 0755)
		if err != nil {
			log.Fatal("Could Not create Cache folder")
		}
	}
	// Check Code.exe
	if _, err := os.Stat(path.Join(CODE_HOME, "MainApp", "Code.exe")); os.IsNotExist(err) {
		err = deployVSCode()
		if err != nil {
			log.Fatal("Could not deploy vscode, ", err)
		}
	}
	// Check ./MainData/tmp
	data_tmpDirPath := path.Join(CODE_HOME, "MainData", "tmp")
	if _, err := os.Stat(data_tmpDirPath); os.IsNotExist(err) {
		err = os.MkdirAll(data_tmpDirPath, 0755)
		if err != nil {
			log.Fatal("Could Not create data folder")
		}
	}
	// Check data link
	dataLinkPath := path.Join(CODE_HOME, "MainApp", "data")
	codeDataFileInfo, err := os.Lstat(dataLinkPath)
	if err != nil {
		err = createDataLink()
		if err != nil {
			log.Fatal("Could not create data link")
		} else {
			codeDataFileInfo, _ = os.Lstat(dataLinkPath)
		}
	}
	if codeDataFileInfo.Mode()&os.ModeSymlink != 0 {
		linkPath, err := os.Readlink(dataLinkPath)
		if err != nil {
			log.Fatal("Could not find data in MainApp folder")
		}
		if linkPath != "..\\MainData" {
			log.Fatal("Data link is not correct")
		}
	}
	// Check Config/code-mint.json
	configFile := path.Join(CODE_HOME, "Config", "code-mint.json")
	if _, err := os.Stat(path.Join(CODE_HOME, "Config")); os.IsNotExist(err) {
		err = os.Mkdir(path.Join(CODE_HOME, "Config"), 0755)
		if err != nil {
			log.Fatal("Could not create Config folder")
		}
	}
	if _, err := os.Stat(path.Join(CODE_HOME, "Config", "code-mint.json")); os.IsNotExist(err) {
		conf, err := os.Create(configFile)
		if err != nil {
			log.Fatal("Could not create config file")
		}
		conf.Write([]byte(configJsonTemplate))
	}

	// Check path include CN character
	if containsChinese(CODE_HOME) {
		log.Fatal("Code Home Path include Chinese character")
	}

}
func containsChinese(path string) bool {
	for _, char := range path {
		if unicode.Is(unicode.Han, char) {
			return true
		}
	}
	return false
}

func deployVSCode() error {
	app_home := path.Join(CODE_HOME, "MainApp")
	err := os.RemoveAll(app_home)
	if err != nil {
		return err
	}
	err = os.Mkdir(app_home, 0755)
	if err != nil {
		return err
	}
	vsc_zipFile := path.Join(CODE_HOME, "Cache", "vscode-x64.zip")
	err = utils.DownloadFile(vsc_zipFile, "https://update.code.visualstudio.com/latest/win32-x64-archive/stable")
	if err != nil {
		return err
	}
	err = utils.UnzipFile(vsc_zipFile, app_home)
	if err != nil {
		return err
	}
	err = os.Remove(vsc_zipFile)
	if err != nil {
		return err
	}
	return nil
}
func createDataLink() error {
	source := "..\\MainData"
	target := "MainApp\\data"
	link := exec.Command("cmd", "/c", "mklink", "/d", target, source)
	link.Dir = CODE_HOME
	err := link.Run()
	if err != nil {
		return err
	}
	return nil
}
