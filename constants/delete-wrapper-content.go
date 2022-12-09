package constants

import (
	"github.com/nsrvel/go-tools/utils"
)

func DeleteWrapperImportContent(file string, layer string, gomod string, category string, domainLower string) string {

	var result string

	if layer == "repo" {
		result = utils.DeleleLineContainSubstring(file, gomod+"/internal/"+category+"/"+domainLower+"/repository")
	} else if layer == "uc" {
		result = utils.DeleleLineContainSubstring(file, gomod+"/internal/"+category+"/"+domainLower+"/usecase")
	} else if layer == "handler" {
		result = utils.DeleleLineContainSubstring(file, gomod+"/internal/"+category+"/"+domainLower+"/delivery")
	} else {
		return ""
	}

	return result
}

func DeleteWrapperStructContent(file string, layer string, category string, domainLower string, domainCap string) string {

	var result string

	if layer == "repo" {
		result = utils.DeleleLineContainSubstring(file, domainLower+"."+"Repository")
	} else if layer == "uc" {
		result = utils.DeleleLineContainSubstring(file, domainLower+"."+"Usecase")
	} else if layer == "handler" {
		result = utils.DeleleLineContainSubstring(file, domainLower+"."+domainCap+"Handler")
	} else {
		return ""
	}

	return result
}

func DeleteWrapperFunctionContent(file string, layer string, category string, domainLower string, domainCap string) string {

	var result string

	if layer == "repo" {
		result = utils.DeleleLineContainSubstring(file, domainLower+".New"+domainCap+"Repo")
	} else if layer == "uc" {
		result = utils.DeleleLineContainSubstring(file, domainLower+".New"+domainCap+"Usecase")
	} else if layer == "handler" {
		result = utils.DeleleLineContainSubstring(file, domainLower+".New"+domainCap+"Handler")
	} else {
		return ""
	}

	return result
}
