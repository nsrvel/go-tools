package views

import (
	"fmt"

	"github.com/nsrvel/go-tools/models"
)

func DisplayFeature(listFeature []models.Feature) {
	fmt.Println("========================================")
	fmt.Println("ID  | Feature")
	fmt.Println("    |               ")
	for _, attr := range listFeature {
		if attr.ID > 0 {
			fmt.Println(fmt.Sprintf("%v   | %s", attr.ID, attr.Name))
		}
		if attr.ID >= 10 {
			fmt.Println(fmt.Sprintf("%v  | %s", attr.ID, attr.Name))

		} else if attr.ID >= 100 {
			fmt.Println(fmt.Sprintf("%v | %s", attr.ID, attr.Name))
		}
	}
	fmt.Println("    |               ")
	fmt.Println("========================================")
}
