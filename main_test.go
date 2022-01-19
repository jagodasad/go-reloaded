package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

// func TestGoReloadedHexToDec(t *testing.T) {

// 	input := "Simply add 42 (hex) and 10 (bin) and you will see the result is 68."
// 	// real expected := "Simply add 66 and 2 and you will see the result is 68."
// 	expected := "Simply add 66 and 10 (bin) and you will see the result is 68."

// 	got := GoReloaded(input)

// 	if expected != got {
// 		t.Errorf("got %q, expected %q", got, expected)
// 	}
// }

// This test file tests the go-reloaded project against the test cases on audit page
func TestGoReloaded(t *testing.T) {
	inputFile, outputFile := "sample.txt", "result.txt"

	/*	Each element of testCases contains a pair of strings, the first string of the
		pair is what is to be written to the input file, the second is what should be
		the contents of the output file */
	testCases := [][]string{
		{
			"If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			"If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?",
		},
		{
			"I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			"I have to pack 5 outfits. Packed 26 just to be sure",
		},
		{
			"Don't be sad ,because sad backwards is das . And das not good",
			"Don't be sad, because sad backwards is das. And das not good",
		},
		{
			"harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			"Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'",
		},
	}

	for _, testCase := range testCases {
		// Write the string to be processed to the input file
		if err := ioutil.WriteFile(inputFile, []byte(testCase[0]), os.ModePerm); err != nil {
			panic(err)
		}

		/*	Attempt to run the main project file with the input and output file
			names as arguments	*/
		if err := exec.Command("go", "run", ".", inputFile, outputFile).Run(); err != nil {
			panic(err)
		}

		/*	Read the contents of the output file, checking if it is equal to the
			expected output */
		if result, err := ioutil.ReadFile(outputFile); err != nil {
			panic(err)
		} else if string(result) != testCase[1] {
			t.Errorf("\nTest fails when given the test case:\n\t\"%s\","+
				"\n%s should contain:\n\t\"%s\",\n%s contains:\n\t\"%s\"\n\n",
				testCase[0], outputFile, testCase[1], outputFile, string(result))
		}
	}
}
