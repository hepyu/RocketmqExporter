package utils

func contains(stringArray []string, str string) int {
	var index = -1
	var length = len(stringArray)
	for i = 0; i < length; i++ {
		if stringArray[i] == str {
			index = i
		}
	}
	return index
}
