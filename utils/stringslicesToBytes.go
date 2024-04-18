package utils

// Function for changing the slices of strings into bytes;
// because os.WriteFile takes bytes and not strings as arguments
func SlicesstringsToBytes(strs []string) []byte {
	var byteslice []byte
	for _, str := range strs {
		str = str + " "
		byteslice = append(byteslice, []byte(str)...)
	}
	return byteslice
}