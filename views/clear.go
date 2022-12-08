package views

import (
	"runtime"

	"github.com/nsrvel/go-tools/utils"
)

func DisplayClear() {
	switch runtime.GOOS {
	case "darwin":
		utils.RunCmd("clear")
	case "linux":
		utils.RunCmd("clear")
	case "windows":
		utils.RunCmd("cmd", "/c", "cls")
	default:
		utils.RunCmd("clear")
	}
}
