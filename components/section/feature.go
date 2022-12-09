package section

import (
	"log"

	"github.com/nsrvel/go-tools/constants"
	"github.com/nsrvel/go-tools/data"
	"github.com/nsrvel/go-tools/utils"
	"github.com/nsrvel/go-tools/views"
)

func FeatureSection(workDir string, gomod string) string {

	views.DisplayClear()

reset:
	listFeature, err := data.ListFeature()
	if err != nil {
		log.Fatalln("program error, failed to get list feature, err: " + err.Error())
	}

	views.DisplayTitle()
	views.DisplayFeature(*listFeature)

	featureID, err := utils.InputAfterQuestion("Select feature id: ")
	if err != nil {
		log.Fatalln(err)
	}

	isFeatureID := utils.CheckIsValidFeatureID(*listFeature, featureID)
	if !isFeatureID {
		views.DisplayResetWithMessage("please select feature id that is already on the list of feature", constants.ColorRed)
		goto reset
	}

	if featureID == "3" {
		views.DisplayResetWithMessage(" sorry this feature under develpoment", constants.ColorYellow)
		goto reset
	}

	return featureID
}
