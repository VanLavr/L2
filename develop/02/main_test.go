package main

import "testing"

func TestUnpackString(t *testing.T) {
	testCases := []struct {
		toUnpack       string
		expectedResult string
		expectedError  error
	}{
		{
			toUnpack:       "a4bc2d5e",
			expectedResult: "aaaabccddddde",
			expectedError:  nil,
		},
		{
			toUnpack:       "a4bc2d5e3",
			expectedResult: "aaaabccdddddeee",
			expectedError:  nil,
		},
		{
			toUnpack:       "12ers",
			expectedResult: "",
			expectedError:  errInvalidString,
		},
		{
			toUnpack:       "abcd",
			expectedResult: "abcd",
			expectedError:  nil,
		},
		{
			toUnpack:       "45",
			expectedResult: "",
			expectedError:  errInvalidString,
		},
		{
			toUnpack:       "",
			expectedResult: "",
			expectedError:  nil,
		},
	}

	for _, tc := range testCases {
		got, err := UnpackString(tc.toUnpack)
		if got != tc.expectedResult || err != tc.expectedError {
			t.Errorf("source string: %s, expected result: %s, expected error: %v\ngot result: %s, got error: %v", tc.toUnpack, tc.expectedResult, tc.expectedError, got, err)
		}
	}
}
