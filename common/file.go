// Package common contains code, that may be reused in every day solutions
package common

import "io/ioutil"

// ReadFileToString reads the file, found on path filename and reads it all into a string.
func ReadFileToString(filename string) string {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("Could not read file " + filename)
	}
	return string(buffer)
}
