package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Please enter the height of your tree:")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input")
		panic(err)
	}

	height, err := ValidateInput(strings.TrimSuffix(input, "\r\n"))
	if err != nil {
		fmt.Println(err)
		panic("err")
	}

	bottomBranch := GetBottomBranchSize(height)

	tree := BuildTree(height, bottomBranch)

	fmt.Println(tree)

}

func BuildTree(height int, bottomBranch int) string {
	var sb strings.Builder
	//Build each row of the tree
	for i := 1; i <= height; i++ {

		//prefix the spaces based on bottom branch width
		numSpaces := (height - i)
		for i := 0; i < numSpaces; i++ {
			_, _ = sb.WriteString(" ")
		}

		// Add the X for each row
		numXs := bottomBranch - (numSpaces * 2)
		for i := 1; i <= numXs; i++ {
			_, _ = sb.WriteString("X")

		}
		_, _ = sb.WriteString("\n")
	}

	//Build base
	numSpaces := (height - 1)
	for i := 0; i < numSpaces; i++ {
		_, _ = sb.WriteString(" ")
	}

	_, _ = sb.WriteString("|")

	return sb.String()
}

func ValidateInput(input string) (int, error) {
	inputHeight, err := strconv.Atoi(input)
	if err != nil || inputHeight <= 1 {
		return 0, errors.New("Error validating input. Please try again with a number > 1")
	}
	return inputHeight, nil
}

func GetBottomBranchSize(height int) (width int) {

	return (height * 2) - 1
}
