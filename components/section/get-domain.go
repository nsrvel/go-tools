package section

import (
	"fmt"
	"log"
	"strings"

	"github.com/nsrvel/go-tools/constants"
	"github.com/nsrvel/go-tools/utils"
	"github.com/nsrvel/go-tools/views"
)

func GetDomainSection(workdir string, category string) string {

	views.DisplayClear()

reset:
	views.DisplayTitle()

	domain, err := utils.InputAfterQuestion("Enter new domain name: ")
	if err != nil {
		log.Fatalln(err)
	}
	err = utils.CheckIsValidDomain(domain)
	if err != nil {
		views.DisplayResetWithMessage(err.Error(), constants.ColorYellow)
		goto reset
	}
	isDomainExist, _ := utils.CheckPathIfExist(workdir + "/internal/" + category + "/" + domain)
	if isDomainExist {
		views.DisplayResetWithMessage("domain already exist", constants.ColorYellow)
		goto reset
	}

reset2:
	fmt.Println("")
	confirm, err := utils.InputAfterQuestion(constants.ColorYellow + "Input ( y ) for next, or ( n ) for back to previous steps : " + constants.ColorReset)
	if err != nil {
		log.Fatalln(err)
	}
	if strings.ToLower(confirm) != "n" && strings.ToLower(confirm) != "y" {
		views.DisplayClear()
		views.DisplayTitle()
		fmt.Println(("Enter new domain name: " + domain))
		goto reset2
	} else if strings.ToLower(confirm) == "n" {
		views.DisplayClear()
		goto reset
	}

	return domain
}
