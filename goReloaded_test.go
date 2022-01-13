package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

// This test file tests go-reloaded project against the test cases on audit page

func TestGoReloaded(t *testing.T) {
	inputFile, outputFile := "sample.txt", "result.txt"


/* Each element of testCases contains a pair of strings, the first string of the pair is what is to be written to the input file, the second is what should be the contents of the output file */

	testCases := [][]string{
		{	"Fear is THE MIND KILLER (low, ) fear is the little death that brings total obliteration. i (cap) will face my fear (up, 2)."

			"Fear is the mind killer fear is the little death that brings total obliteration. I will face MY FEAR."
		},

		{
			"We will complete 1010000 (bin) projects. Then we will land our 01 (hex) st IT jobs",

			"We will complete 50 projects. Then we will land our 01 st IT jobs"
		},
		{
			"I went to Roma, because Roma backwards is amoR. And amoR is good",

			"I went to Roma, because Roma backwards is amoR. And amoR is good",
		}

		{
			"jackie kennedy: (cap, 2) : ' Pearls are always appropriate ;diamonds are appropriate after 6pm .'",
			"Jackie Kennedy: 'Pearls are always appropriate; diamonds are appropriate after 6pm.'"
		},
	}

	for _, testCase := range testCases {
		//Write the string to be processed to the input file
		if err := ioutil.WriteFile(inputFile, []byte(testCase[0]), os.ModePerm); err != nil {
			panic(err)
		}
		/* Attempt to run the main project file with the input and output file names as arguments */
		if err := exec.Command("go", "run", ".", inputFile, outputFile).Run(); err != nil {
			panic(err)
		}
		/* Read the contents of the output file, checking if it is equal to the expected output */
		if result, err := ioutil.ReadFile(outputFile); err != nil {
			panic(err)
		} else if string(result) != testCase[1] {
			t.Error("\nTest fails when given the test case:\n\t\"%s\","+
					"\n%s should contain:\n\t\"%s\",\n%s contains:\n\t\"%s\"\n\n",
				testCase[0], outputFile, testCase[1], outputFile, string(result))
		}
	}
}