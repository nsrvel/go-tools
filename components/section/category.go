package section

import (
	"log"

	"github.com/nsrvel/go-tools/constants"
	"github.com/nsrvel/go-tools/utils"
	"github.com/nsrvel/go-tools/views"
)

func CategorySection() string {

	views.DisplayClear()

reset:
	views.DisplayTitle()
	views.DisplayCategory()

	category, err := utils.InputAfterQuestion("Select category id: ")
	if err != nil {
		log.Fatalln(err)
	}

	if category == "1" {
		category = "general"
	} else if category == "2" {
		category = "core"
	} else if category == "3" {
		category = "cms"
	} else {
		views.DisplayResetWithMessage("please select category id that is already on the list of category", constants.ColorRed)
		goto reset
	}
	return category
}
