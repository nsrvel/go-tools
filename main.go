package main

import (
	"log"
	"os"

	createdomain "github.com/nsrvel/go-tools/components/create-domain"
	deletedomain "github.com/nsrvel/go-tools/components/delete-domain"
	"github.com/nsrvel/go-tools/components/section"
	"github.com/nsrvel/go-tools/constants"
	"github.com/nsrvel/go-tools/utils"
	"github.com/nsrvel/go-tools/views"
)

func main() {

	var err error

	//* Get working directory
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalln("program error, failed to get current working directory")
	}

	//* Get go mod name
	gomod, err := utils.GetGoModName(workDir)
	if err != nil {
		log.Fatalln("program error, " + err.Error())
	}

	//* Display section feature
	featureID := section.FeatureSection(workDir, gomod)

	//* Run function depend on feature id
	if featureID == "1" {
		err = createdomain.CreateDomain(workDir, gomod)
	} else if featureID == "2" {
		err = deletedomain.DeleteDomain(workDir, gomod)
	}
	if err != nil {
		views.DisplayResetWithMessage("program error, "+err.Error(), constants.ColorRed)
	}
}
