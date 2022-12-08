package views

import (
	"fmt"

	"github.com/nsrvel/go-tools/constants"
)

func DisplayTitle() {
	fmt.Println("")
	fmt.Println(constants.ColorCyan, "Go-Tools v1.0.1", constants.ColorReset)
	fmt.Println("")
}
