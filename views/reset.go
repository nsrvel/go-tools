package views

import (
	"fmt"

	"github.com/nsrvel/go-tools/constants"
)

func DisplayResetWithMessage(message string, customColor string) {
	DisplayClear()
	if customColor != "" {
		fmt.Println("message: ", customColor, message, constants.ColorReset)
	} else {
		fmt.Println("message: ", constants.ColorGreen, message, constants.ColorReset)
	}
	fmt.Println("")
}
