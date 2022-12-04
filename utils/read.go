package utils

// ReadExample will attempt to read the expected example file with the expected name of ex.txt located in the day's package folder
func ReadExample() []string {
	return ReadLinesFromFile("ex.txt")
}

// ReadInput will attempt to read the expected input file with the expected name of in.txt located in the day's package folder
func ReadInput() []string {
	return ReadLinesFromFile("in.txt")
}
