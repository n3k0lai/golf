package lib

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func AuditConfig(config ChatterinoConfig) bool {
	AuditHighlights(config)
	return true
}
func removeDuplicateHighlights(intSlice []IndividualHighlight) []IndividualHighlight {
	keys := make(map[string]bool)
	list := []IndividualHighlight{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry.Pattern]; !value {
			keys[entry.Pattern] = true
			list = append(list, entry)
		}
	}
	return list
}
func AuditHighlights(config ChatterinoConfig) bool {
	println("Auditing " + fmt.Sprint(len(config.Highlighting.Users)) + " user highlights")
	// iterate through highlights
	visited := make(map[string]int, 0)
	for i := 0; i < len(config.Highlighting.Users); i++ {
		visited[config.Highlighting.Users[i].Pattern]++
	}

	// if there are any duplicates, note their index
	duplicatesFound := false
	for key, element := range visited {
		if element > 1 {
			if !duplicatesFound {
				duplicatesFound = true

				fmt.Println("Duplicate Keys found: ")
			}
			fmt.Print("* ")
			color.Red(fmt.Sprintf("%s => Count: %s", key, strconv.Itoa(element)))

		}
	}

	if duplicatesFound {
		//println(len(visited))
		validate := func(input string) error {
			if strings.ToLower(input) == "y" || strings.ToLower(input) == "n" || input == "" {
				return nil
			}
			//if err != nil {
			return errors.New("invalid input")
			//}
			//return nil
		}

		prompt := promptui.Prompt{
			Label:    "Do you wish to prune duplicate highlights? (Y/n)",
			Validate: validate,
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return false
		}

		if result == "y" || result == "" {
			fmt.Println("Pruning...")
			config.Highlighting.Users = removeDuplicateHighlights(config.Highlighting.Users)
		}
	}

	return true
}
