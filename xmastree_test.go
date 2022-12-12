package main

import (
	"errors"
	"testing"

	"github.com/go-playground/assert"
)

func TestGetBottomBranchLength(t *testing.T) {
	testCases := []struct {
		treeHeight     int
		expectedOutput int
	}{
		{
			treeHeight:     2,
			expectedOutput: 3,
		},
		{
			treeHeight:     3,
			expectedOutput: 5,
		},
		{
			treeHeight:     4,
			expectedOutput: 7,
		},
		{
			treeHeight:     5,
			expectedOutput: 9,
		},
		{
			treeHeight:     6,
			expectedOutput: 11,
		},
	}

	for _, tC := range testCases {
		output := GetBottomBranchSize(tC.treeHeight)

		assert.Equal(t, output, tC.expectedOutput)
	}
}

func TestValidateInput(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput int
		expectedError  error
	}{
		{
			input:          "2",
			expectedOutput: 2,
		},
		{
			input:          "10",
			expectedOutput: 10,
		},
		{
			input:          "string",
			expectedOutput: 0,
			expectedError:  errors.New("Error validating input. Please try again with a number > 1"),
		},
		{
			input:          "0",
			expectedOutput: 0,
			expectedError:  errors.New("Error validating input. Please try again with a number > 1"),
		},
		{
			input:          "1",
			expectedOutput: 0,
			expectedError:  errors.New("Error validating input. Please try again with a number > 1"),
		},
	}

	for _, tC := range testCases {
		output, err := ValidateInput(tC.input)

		assert.Equal(t, output, tC.expectedOutput)
		assert.Equal(t, err, tC.expectedError)
	}
}

func TestBuildTree(t *testing.T) {
	testCases := []struct {
		treeHeight        int
		bottomBranchWidth int
		expectedOutput    string
	}{
		{
			treeHeight:        2,
			bottomBranchWidth: 3,
			expectedOutput:    " X\nXXX\n |",
		},
		{
			treeHeight:        3,
			bottomBranchWidth: 5,
			expectedOutput:    "  X\n XXX\nXXXXX\n  |",
		},
		{
			treeHeight:        4,
			bottomBranchWidth: 7,
			expectedOutput:    "   X\n  XXX\n XXXXX\nXXXXXXX\n   |",
		},
	}

	for _, tC := range testCases {
		output := BuildTree(tC.treeHeight, tC.bottomBranchWidth)

		assert.Equal(t, output, tC.expectedOutput)
	}
}
