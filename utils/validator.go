package utils

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/nsrvel/go-tools/models"
)

func CheckIsValidFeatureID(listFeature []models.Feature, id string) bool {
	isOK := false
	for _, attr := range listFeature {
		if fmt.Sprintf("%v", attr.ID) == id {
			isOK = true
		}
	}
	return isOK
}

func CheckPathIfExist(path string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

func CheckPathIfNotExist(path string) (bool, error) {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

func CheckIsValidDomain(domain string) error {

	length := len(domain)
	firstChar := domain[0:1]

	rule1 := regexp.MustCompile(`^[a-zA-Z_]*$`)
	if !rule1.MatchString(domain) {
		return errors.New("domain can contain only alphabets and underscores ( _ )")
	}
	rule2 := regexp.MustCompile(`^[a-z]*$`)
	if !rule2.MatchString(firstChar) {
		return errors.New("domain cannot start with capital character or a symbol")
	}
	if strings.Contains(domain, "__") {
		return errors.New("domain cannot contain consecutive symbols")
	}
	min := 2
	if length < min {
		return errors.New(fmt.Sprintf("domain cannot be less than %v characters", min))
	}
	max := 30
	if length < min {
		return errors.New(fmt.Sprintf("domain cannot be less than %v characters", max))
	}
	return nil
}
